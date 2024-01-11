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

// ToDo: implement ReservedSeats in its sql query
func (a *API) GetTicket(ctx iris.Context) {
	claims, _ := ctx.Values().Get("claims").(*Claims)

	tid, _ := ctx.Params().GetInt64("id")

	ticket, err := sql.SelectTicket(ctx, a.db, tid, claims.UserID)
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

	tickets, err := sql.SelectUserTickets(ctx, a.db, claims.UserID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, NewError("Failed to fetch your tickets, please try again later.", err))
		return
	}

	ctx.JSON(GetTicketsResponse{Tickets: tickets})
}

func (a *API) PostReserveTickets(ctx iris.Context) { ctx.StatusCode(iris.StatusNotImplemented) }
