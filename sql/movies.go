package sql

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type MovieSummary struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ShowTime int16  `db:"show_time" json:"showTime"`
}

func SelectMovies(ctx context.Context, db Querier) ([]MovieSummary, error) {
	movies := []MovieSummary{}
	err := pgxscan.Select(ctx, db, &movies, `select id, name, show_time from movies order by show_time asc`)
	return movies, err
}

type Movie struct {
	ID            int64
	Name          string
	Time          int16
	Genres        []string
	ReservedSeats []int16
}
