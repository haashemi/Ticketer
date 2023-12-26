package main

import (
	"github.com/haashemi/Ticketer/api"
	"github.com/haashemi/Ticketer/internal/config"
	"github.com/haashemi/Ticketer/sql"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	conn, err := sql.NewConnection(conf.Database)
	if err != nil {
		panic(err)
	}

	if err = sql.Migrate(conf.Database); err != nil {
		panic(err)
	}

	api.Run(conf, conn)
}
