package api

import (
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
)

func (a *API) GetMovies(ctx iris.Context) {
	movies, err := sql.SelectMovies(ctx, a.db)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch movies list, please try again later.", err))
		return
	}

	ctx.JSON(map[string]any{"movies": movies})
}

func (a *API) GetMovie(ctx iris.Context) {

}
