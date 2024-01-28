package api

import (
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
	"github.com/lib/pq"
)

func (a *API) GetMovies(ctx iris.Context) {
	movies, err := sql.SelectMovies(a.db)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch movies list, please try again later.", err))
		return
	}

	ctx.JSON(map[string]any{"movies": movies})
}

func (a *API) GetMovie(ctx iris.Context) {
	mid, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, NewError("Request is not valid", err))
		return
	}

	movie, err := sql.SelectMovie(a.db, mid)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch movie, please try again later.", err))
		return
	}

	ctx.JSON(movie)
}

type GetMovieReservedSeatsResponse struct {
	ReservedSeats pq.Int32Array `json:"reservedSeats"`
}

func (a *API) GetMovieReservedSeats(ctx iris.Context) {
	mid, err := ctx.Params().GetInt64("id")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, NewError("Request is not valid", err))
		return
	}

	date, err := ctx.Params().GetTime("date")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, NewError("Request is not valid", err))
		return
	}

	reservedSeats, err := sql.SelectMovieReservedSeats(a.db, mid, date)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch reserved seats, please try again later.", err))
		return
	}

	ctx.JSON(GetMovieReservedSeatsResponse{ReservedSeats: reservedSeats})
}
