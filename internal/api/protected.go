package api

import (
	"time"

	"github.com/haashemi/Ticketer/internal/postgres"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kataras/iris/v12"
)

type GetProfileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *API) GetProfile(ctx iris.Context) {
	claims := ctx.Values().Get("claims").(*Claims)

	u, err := a.db.SelectUser(ctx, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch profile", err))
		return
	}

	ctx.JSON(GetProfileResponse{Name: u.FullName, Email: u.Email})
}

// ToDo: implement ReservedSeats in its sql query
func (a *API) GetTicket(ctx iris.Context) {
	claims, _ := ctx.Values().Get("claims").(*Claims)

	tid, _ := ctx.Params().GetInt64("id")

	ticket, err := a.db.SelectTicket(ctx, postgres.SelectTicketParams{ID: tid, UserID: claims.UserID})
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch ticket info, please try again later.", err))
		return
	}

	ctx.JSON(ticket)
}

type GetTicketsResponse struct {
	Tickets []postgres.SelectUserTicketsRow `json:"tickets"`
}

func (a *API) GetTickets(ctx iris.Context) {
	claims, _ := ctx.Values().Get("claims").(*Claims)

	tickets, err := a.db.SelectUserTickets(ctx, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch your tickets, please try again later.", err))
		return
	}

	ctx.JSON(GetTicketsResponse{Tickets: tickets})
}

type PostReserveTicketsRequest struct {
	MovieID  int64   `json:"movieId" validate:"required"`
	ForYear  int     `json:"forYear" validate:"required"`
	ForMonth int     `json:"forMonth" validate:"required"`
	ForDay   int     `json:"forDay" validate:"required"`
	Seats    []int16 `json:"seats" validate:"required,min=0"`
}

type PostReserveTicketsResponse struct {
	TicketID int64 `json:"ticketId"`
}

func (a *API) PostReserveTickets(ctx iris.Context) {
	var body PostReserveTicketsRequest
	if err := ctx.ReadJSON(&body); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, NewError("Request is not valid", err))
		return
	}

	claims, _ := ctx.Values().Get("claims").(*Claims)

	// Create a DB transaction
	dbTx, err := a.db.Conn.Begin(ctx)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Database error.", err))
		return
	}
	defer dbTx.Rollback(ctx)

	tx := a.db.WithTx(dbTx)

	movie, err := tx.SelectMovie(ctx, body.MovieID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to find the movie", err))
		return
	}

	uu := time.Date(body.ForYear, time.Month(body.ForMonth), body.ForDay, 0, 0, 0, 0, time.UTC)

	tid, err := tx.InsertTicket(ctx, postgres.InsertTicketParams{
		UserID:       claims.UserID,
		MovieID:      movie.ID,
		PremiereDate: pgtype.Date{Time: uu, Valid: true},
		PremiereTime: movie.PremiereTime,
	})
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to get you a ticket", err))
		return
	}

	for _, seat := range body.Seats {
		err = tx.InsertSeat(ctx, postgres.InsertSeatParams{TicketID: tid, SeatNumber: seat})
		if err != nil {
			ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to reserve your seat", err))
			return
		}
	}

	err = dbTx.Commit(ctx.Request().Context())
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to submit the data", err))
		return
	}

	ctx.JSON(PostReserveTicketsResponse{TicketID: tid})
}
