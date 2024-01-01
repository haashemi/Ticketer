package sql

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type MovieSummary struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	PremiereTime time.Time `db:"premiere_time" json:"premiereTime"`
}

func SelectMovies(ctx context.Context, db Querier) ([]MovieSummary, error) {
	movies := []MovieSummary{}
	err := pgxscan.Select(ctx, db, &movies, `select id, name, premiere_time from movies order by show_time asc`)
	return movies, err
}

type Movie struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	MovieTime    int16     `db:"movie_time" json:"movieTime"`
	Genres       []string  `db:"genres" json:"genres"`
	FromDate     time.Time `db:"from_date" json:"fromDate"`
	ToDate       time.Time `db:"to_date" json:"toDate"`
	PremiereTime time.Time `db:"premiere_time" json:"premiereTime"`
}

func SelectMovie(ctx context.Context, db Querier, id int64) (data Movie, err error) {
	err = pgxscan.Get(ctx, db, &data, `select id, name, movie_time, genres, from_date, to_date, premiere_time from movies where id = $1`, id)
	return
}
