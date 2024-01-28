package api

import (
	"time"

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

	u, err := sql.SelectUser(a.db, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch profile", err))
		return
	}

	ctx.JSON(GetProfileResponse{Name: u.FullName, Email: u.Email, IsAdmin: u.IsAdmin})
}

// ToDo: implement ReservedSeats in its sql query
func (a *API) GetTicket(ctx iris.Context) {
	claims, _ := ctx.Values().Get("claims").(*Claims)

	tid, _ := ctx.Params().GetInt64("id")

	ticket, err := sql.SelectTicket(a.db, tid, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch ticket info, please try again later.", err))
		return
	}

	ctx.JSON(ticket)
}

type GetTicketsResponse struct {
	Tickets []sql.TicketSummary `json:"tickets"`
}

func (a *API) GetTickets(ctx iris.Context) {
	claims, _ := ctx.Values().Get("claims").(*Claims)

	tickets, err := sql.SelectUserTickets(a.db, claims.UserID)
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
	Seats    []uint8 `json:"seats" validate:"required"`
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

	var tid int64
	err := a.db.Transaction(func(tx sql.Queryer) error {
		movie, err := sql.SelectMovie(tx, body.MovieID)
		if err != nil {
			return NewError("Failed to find the movie", err)
		}

		uu := time.Date(body.ForYear, time.Month(body.ForMonth), body.ForDay, 0, 0, 0, 0, time.UTC)

		tid, err = sql.InsertTicket(tx, claims.UserID, movie.ID, uu, movie.PremiereTime)
		if err != nil {
			return NewError("Failed to get you a ticket", err)
		}

		for _, seat := range body.Seats {
			err = sql.InsertSeat(tx, tid, seat)
			if err != nil {
				return NewError("Failed to reserve your seat", err)
			}
		}

		return nil
	})

	if err != nil {
		if _, ok := err.(Error); ok {
			ctx.StopWithJSON(iris.StatusInternalServerError, err)
		} else {
			ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to reserve, please try again later.", err))
		}
		return
	}

	ctx.JSON(PostReserveTicketsResponse{TicketID: tid})
}
