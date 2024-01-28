package sql

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type MovieSummary struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	PremiereTime time.Time `db:"premiere_time" json:"premiereTime"`
}

func SelectMovies(db Queryer) ([]MovieSummary, error) {
	movies := []MovieSummary{}
	err := sqlx.Select(db, &movies, `select id, name, premiere_time from movies order by premiere_time asc`)
	return movies, err
}

type Movie struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	MovieTime    int16          `db:"movie_time" json:"movieTime"`
	Genres       pq.StringArray `db:"genres" json:"genres"`
	FromDate     time.Time      `db:"from_date" json:"fromDate"`
	ToDate       time.Time      `db:"to_date" json:"toDate"`
	PremiereTime time.Time      `db:"premiere_time" json:"premiereTime"`
}

func SelectMovie(db Queryer, id int64) (data Movie, err error) {
	err = sqlx.Get(db, &data, `select id, name, movie_time, genres, from_date, to_date, premiere_time from movies where id = $1`, id)
	return
}

func SelectMovieReservedSeats(db Queryer, id int64, date time.Time) (data pq.Int32Array, err error) {
	data = pq.Int32Array{}
	err = sqlx.Select(db, &data, `
		select 
			s.seat_number 
		from 
			tickets t 
			left join seats s on s.ticket_id = t.id
		where
			t.movie_id = $1
			and t.premiere_date = $2
	`, id, date)
	return
}
