package api

import (
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
)

const TokenTime = 7 * 24 * time.Hour

var (
	ErrClaimsNotInContext = errors.New("claims not found in the context")
	ErrInvalidClaims      = errors.New("invalid claims")
	ErrTokenExpired       = errors.New("token is expired")
	ErrNotAdmin           = errors.New("you are not an admin")
)

type Claims struct {
	UserID  int64 `json:"id"`
	IsAdmin bool  `json:"isAdmin"`
	jwt.RegisteredClaims
}

func (a *API) newToken(userID int64, isAdmin bool) (t string, err error) {
	currentTime := time.Now()

	claims := &Claims{
		UserID:  userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.conf.Host,
			Audience:  []string{a.conf.Host},
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(TokenTime)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.conf.JWTKey))
}

func (a *API) setAccessTokenCookie(ctx iris.Context, token string) {
	ctx.SetCookieKV("accessToken", token, iris.CookieExpires(TokenTime))
}

type SignInBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

func (a *API) SignIn(ctx iris.Context) {
	var body SignInBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	user, err := sql.SelectUserEssentials(ctx, a.db, body.Email)
	if err != nil {
		if err.Error() == "" {
			ctx.StopWithText(iris.StatusForbidden, "Invalid username or password")
			return
		}
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	hashedPassword, err := hex.DecodeString(user.Password)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(body.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			ctx.StopWithText(iris.StatusForbidden, "Invalid username or password")
			return
		}
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	ctx.JSON(map[string]string{"name": user.FullName})
}

type SignUpBody struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

func (a *API) SignUp(ctx iris.Context) {
	var body SignUpBody
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	id, err := sql.InsertUser(ctx, a.db, body.FullName, body.Email, hex.EncodeToString(password))
	if err != nil {
		// ToDo: find a better way, I have no idea why I'm doing it this way. I'm so sorry. forgive me please.
		if err.Error() == `scanning one: scany: rows final error: ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			ctx.StopWithText(iris.StatusConflict, "Email already exists. try to sign-in or use another email.")
			return
		}
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	token, err := a.newToken(id, false)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	a.setAccessTokenCookie(ctx, token)

	ctx.StatusCode(iris.StatusNoContent)
}

func (a *API) SignOut(ctx iris.Context) {
	// ToDo: set their token to a KV storage as invalidated

	ctx.RemoveCookie("accessToken")
	ctx.StatusCode(iris.StatusNoContent)
}

func (a *API) doCheckAuth(ctx iris.Context) {
	token, err := jwt.ParseWithClaims(
		ctx.GetCookie("accessToken"),
		&Claims{},
		func(t *jwt.Token) (interface{}, error) { return []byte(a.conf.JWTKey), nil },
		jwt.WithIssuer(a.conf.Host),
		jwt.WithAudience(a.conf.Host),
	)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		ctx.StopWithError(iris.StatusBadRequest, ErrInvalidClaims)
		return
	}

	// ToDo: add a KV Storage and store the invalidated tokens.
	// and check if the currentToken is not there.

	// ToDo: make sure that the user is not banned

	ctx.Values().Set("claims", claims)
}

func (a *API) doRefreshToken(ctx iris.Context) {
	claims, ok := ctx.Values().Get("claims").(*Claims)
	if !ok {
		ctx.StopWithError(iris.StatusInternalServerError, ErrClaimsNotInContext)
		return
	}

	currentTime := time.Now()
	claims.ExpiresAt = jwt.NewNumericDate(currentTime.Add(TokenTime))
	claims.IssuedAt = jwt.NewNumericDate(currentTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.conf.JWTKey))
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	a.setAccessTokenCookie(ctx, tokenString)

	ctx.Next()
}

func (a *API) doCheckAdmin(ctx iris.Context) {
	claims, ok := ctx.Values().Get("claims").(*Claims)
	if !ok {
		ctx.StopWithError(iris.StatusInternalServerError, ErrClaimsNotInContext)
		return
	} else if claims.IsAdmin {
		ctx.StopWithError(iris.StatusForbidden, ErrNotAdmin)
		return
	}
	ctx.Next()
}
