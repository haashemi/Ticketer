package api

import (
	"github.com/haashemi/Ticketer/internal/postgres"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kataras/iris/v12"
)

func (a *API) GetMovies(ctx iris.Context) {
	movies, err := a.db.SelectMovies(ctx)
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

	movie, err := a.db.SelectMovie(ctx, mid)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch movie, please try again later.", err))
		return
	}

	ctx.JSON(movie)
}

type GetMovieReservedSeatsResponse struct {
	ReservedSeats []int16 `json:"reservedSeats"`
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

	reservedSeats, err := a.db.SelectMovieReservedSeats(ctx, postgres.SelectMovieReservedSeatsParams{MovieID: mid, PremiereDate: pgtype.Date{Time: date, Valid: true}})
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch reserved seats, please try again later.", err))
		return
	}

	ctx.JSON(GetMovieReservedSeatsResponse{ReservedSeats: reservedSeats})
}
