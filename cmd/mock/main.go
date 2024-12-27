package main

import (
	"context"
	"log"

	"github.com/haashemi/Ticketer/internal/config"
	"github.com/haashemi/Ticketer/internal/mock"
	"github.com/haashemi/Ticketer/internal/postgres"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalln("Failed to load the config", err)
	}

	conn, err := postgres.Connect(conf.Database)
	if err != nil {
		log.Fatalln("Failed to connect to the database", err)
	} else if err = conn.Migrate(); err != nil {
		log.Fatalln("Failed to run the migrations", err)
	}

	err = mock.MockMovies(context.Background(), conn)
	if err != nil {
		log.Fatalln("Failed to mock the movies", err)
	}
}
