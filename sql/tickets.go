package sql

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgtype"
)

type Ticket struct {
	MovieID       int64     `db:"movie_id" json:"movieId"`
	MovieName     string    `db:"movie_name" json:"movieName"`
	MovieTime     int8      `db:"movie_time" json:"movieTime"`
	MovieGenres   []string  `db:"movie_genres" json:"movieGenres"`
	ReservedAt    time.Time `db:"reserved_at" json:"reservedAt"`
	PremiereDate  time.Time `db:"premiere_date" json:"premiereDate"`
	PremiereTime  *Time     `db:"premiere_time" json:"premiereTime"`
	ReservedSeats []int8    `db:"reserved_seats" json:"reservedSeats"`
}

func SelectTicket(ctx context.Context, db Querier, ticketID, userID int64) (data Ticket, err error) {
	err = pgxscan.Get(ctx, db, &data, `
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
	PremiereTime *Time     `db:"premiere_time" json:"premiereTime"`
}

func SelectUserTickets(ctx context.Context, db Querier, user int64) (data []TicketSummary, err error) {
	data = []TicketSummary{}
	err = pgxscan.Select(ctx, db, &data, `
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

func InsertTicket(ctx context.Context, db Querier, uid, mid int64, pDate time.Time, pTime pgtype.Time) (id int64, err error) {
	err = pgxscan.Get(ctx, db, &id, `insert into tickets (user_id, movie_id, premiere_date, premiere_time) values ($1, $2, $3, $4) returning id`, uid, mid, pDate, pTime)
	return
}

func InsertSeat(ctx context.Context, db Querier, tid int64, seat int8) error {
	_, err := db.Exec(ctx, `insert into seats (ticket_id, seat_number) values ($1, $2)`, tid, seat)
	return err
}
