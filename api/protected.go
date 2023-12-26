package api

import (
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
)

func (a *API) GetProfile(ctx iris.Context) {
	claims := ctx.Values().Get("claims").(*Claims)

	profile, err := sql.SelectUser(ctx, a.db, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch profile", err))
		return
	}

	ctx.JSON(profile)
}
