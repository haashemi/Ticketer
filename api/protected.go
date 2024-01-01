package api

import (
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
)

type GetProfileResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

func (a *API) GetProfile(ctx iris.Context) {
	claims := ctx.Values().Get("claims").(*Claims)

	u, err := sql.SelectUser(ctx, a.db, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch profile", err))
		return
	}

	ctx.JSON(GetProfileResponse{Name: u.FullName, Email: u.Email, IsAdmin: u.IsAdmin})
}
