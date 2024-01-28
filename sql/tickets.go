package sql

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Ticket struct {
	MovieID       int64          `db:"movie_id" json:"movieId"`
	MovieName     string         `db:"movie_name" json:"movieName"`
	MovieTime     int8           `db:"movie_time" json:"movieTime"`
	MovieGenres   pq.StringArray `db:"movie_genres" json:"movieGenres"`
	ReservedAt    time.Time      `db:"reserved_at" json:"reservedAt"`
	PremiereDate  time.Time      `db:"premiere_date" json:"premiereDate"`
	PremiereTime  time.Time      `db:"premiere_time" json:"premiereTime"`
	ReservedSeats pq.Int32Array  `db:"reserved_seats" json:"reservedSeats"`
}

func SelectTicket(db Queryer, ticketID, userID int64) (data Ticket, err error) {
	err = sqlx.Get(db, &data, `
		select 
			m.id			as movie_id				,
			m.name 			as movie_name			,
			m.movie_time 	as movie_time			,
			m.genres 		as movie_genres			,
			t.reserved_at 	as reserved_at			,
			t.premiere_date as premiere_date		,
			t.premiere_time as premiere_time		,
			array_agg(s.seat_number) as reserved_seats
		from
			tickets t
			left join seats s on s.ticket_id = t.id
			inner join movies m on m.id = t.movie_id
		where 
			t.id = $1
			and t.user_id = $2
		group by
			t.id, m.id
	 `, ticketID, userID)
	return
}

type TicketSummary struct {
	MovieID      int64     `db:"movie_id" json:"movieID"`
	TicketID     int64     `db:"ticket_id" json:"ticketID"`
	MovieName    string    `db:"movie_name" json:"movieName"`
	PremiereDate time.Time `db:"premiere_date" json:"premiereDate"`
	PremiereTime time.Time `db:"premiere_time" json:"premiereTime"`
}

func SelectUserTickets(db Queryer, user int64) (data []TicketSummary, err error) {
	data = []TicketSummary{}
	err = sqlx.Select(db, &data, `
		select 
			m.id 			as movie_id		,
			t.id 			as ticket_id	,
			m.name 			as movie_name	,
			t.premiere_date as premiere_date,
			t.premiere_time as premiere_time
		from
			tickets t
			join movies m on m.id = t.movie_id
		where 
			t.user_id = $1
		order by
			reserved_at desc
	 `, user)
	return
}

func InsertTicket(db Queryer, uid, mid int64, pDate time.Time, pTime time.Time) (id int64, err error) {
	err = sqlx.Get(db, &id, `insert into tickets (user_id, movie_id, premiere_date, premiere_time) values ($1, $2, $3, $4) returning id`, uid, mid, pDate, pTime)
	return
}

func InsertSeat(db Queryer, tid int64, seat uint8) error {
	_, err := db.Exec(`insert into seats (ticket_id, seat_number) values ($1, $2)`, tid, seat)
	return err
}
