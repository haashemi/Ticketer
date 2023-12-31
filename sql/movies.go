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
	ID     int64    `db:"id" json:"id"`
	Name   string   `db:"name" json:"name"`
	Time   int16    `db:"movie_time" json:"time"`
	Genres []string `db:"genres" json:"genres"`
}

func SelectMovie(ctx context.Context, db Querier, id int64) (data Movie, err error) {
	err = pgxscan.Get(ctx, db, &data, `select id, name, movie_time, genres from movies where id = $1`, id)
	return
}
