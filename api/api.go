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
		auth.Post("/sign-in", api.SignIn) // ToDo
		auth.Post("/sign-up", api.SignUp) // ToDo
		auth.Post("/sign-out", api.SignOut)
	}

	public := app.Party("/api")
	{
		public.Get("/movies", api.GetMovies)
		public.Get("/movies/{id:number}", api.GetMovie)
	}

	protected := app.Party("/api", api.doCheckAuth, api.doRefreshToken)
	{
		protected.Get("/tickets/reserve", NotImplemented)
		protected.Get("/profile", NotImplemented)
		protected.Get("/profile/tickets", NotImplemented)
	}

	adminOnly := app.Party("/api", api.doCheckAuth, api.doRefreshToken, api.doCheckAdmin)
	{
		adminOnly.Post("/movies", NotImplemented)
		adminOnly.Patch("/movies", NotImplemented)
		adminOnly.Delete("/movies", NotImplemented)

		adminOnly.Get("/users", NotImplemented)
		adminOnly.Patch("/users", NotImplemented)
	}

	target, _ := url.Parse("http://localhost:4173")
	proxy := host.ProxyHandler(target, nil)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) { proxy.ServeHTTP(ctx.ResponseWriter(), ctx.Request()) })

	app.Listen(conf.APIAddr)
}

func NotImplemented(ctx iris.Context) { ctx.StatusCode(iris.StatusNotImplemented) }
