package api

import (
	"net/url"

	"github.com/go-playground/validator"
	"github.com/haashemi/Ticketer/internal/config"
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/host"
)

type API struct {
	db   *sql.Connection
	conf *config.Config
}

func Run(conf *config.Config, db *sql.Connection) {
	api := &API{db: db, conf: conf}

	app := iris.Default()
	app.Validator = validator.New()
	app.Use(iris.Compression)

	auth := app.Party("/api/auth")
	{
		auth.Post("/sign-in", api.SignIn)
		auth.Post("/sign-up", api.SignUp)
		auth.Post("/sign-out", api.SignOut)
	}

	public := app.Party("/api/public")
	{
		public.Get("/movies", api.GetMovies)
		public.Get("/movies/{id:number}", api.GetMovie)
		public.Get("/movies/{id:number}/reserved-seats/{date:date}", api.GetMovieReservedSeats)
	}

	protected := app.Party("/api/profile", api.doCheckAuth, api.doRefreshToken)
	{
		protected.Get("/", api.GetProfile)
		protected.Get("/ticket/{id:number}", api.GetTicket)
		protected.Get("/tickets", api.GetTickets)
		protected.Post("/tickets/reserve", api.PostReserveTickets)
	}

	// adminOnly := app.Party("/api/admin", api.doCheckAuth, api.doRefreshToken, api.doCheckAdmin)
	// {
	// 	adminOnly.Post("/movies", NotImplemented)
	// 	adminOnly.Patch("/movies", NotImplemented)
	// 	adminOnly.Delete("/movies", NotImplemented)

	// 	adminOnly.Get("/users", NotImplemented)
	// 	adminOnly.Patch("/users", NotImplemented)
	// }

	app.HandleDir("/static", iris.Dir("./static"))

	target, _ := url.Parse("http://localhost:5173")
	proxy := host.ProxyHandler(target, nil)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) { proxy.ServeHTTP(ctx.ResponseWriter(), ctx.Request()) })

	app.Listen(conf.APIAddr)
}
