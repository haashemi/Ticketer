package api

import (
	"github.com/go-playground/validator"
	"github.com/haashemi/Ticketer/internal/config"
	"github.com/haashemi/Ticketer/sql"
	"github.com/kataras/iris/v12"
)

type API struct {
	db   *sql.Connection
	conf *config.Config
}

func Run(conf *config.Config, db *sql.Connection) {
	// api := &API{db: db, conf: conf}

	app := iris.Default()
	app.Validator = validator.New()
	app.Use(iris.Compression)

	auth := app.Party("/api/auth")
	{
		auth.Post("/login", NotImplemented)
		auth.Post("/logout", NotImplemented)
		auth.Post("/refresh", NotImplemented)
	}

	movies := app.Party("/api/movies")
	{
		movies.Get("/", NotImplemented)
		movies.Get("/{id:number}", NotImplemented)
		movies.Post("/", NotImplemented)
		movies.Patch("/", NotImplemented)
		movies.Delete("/", NotImplemented)
	}

	tickets := app.Party("/api/tickets")
	{
		tickets.Get("/reserve", NotImplemented)
	}

	users := app.Party("/users")
	{
		users.Get("/", NotImplemented)
		users.Patch("/", NotImplemented)
	}

	profile := app.Party("/api/profile")
	{
		profile.Get("/", NotImplemented)
		profile.Get("/tickets", NotImplemented)
	}

	app.Listen(conf.APIAddr)
}

func NotImplemented(ctx iris.Context) { ctx.StatusCode(iris.StatusNotImplemented) }
