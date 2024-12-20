-- name: InsertTicket :one
insert into tickets (user_id, movie_id, premiere_date, premiere_time) values ($1, $2, $3, $4) returning id;

-- name: InsertSeat :exec
insert into seats (ticket_id, seat_number) values ($1, $2);

-- name: SelectTicket :one
select 
    m.id			as movie_id				,
    m.name 			as movie_name			,
    m.movie_time 	as movie_time			,
    m.genres 		as movie_genres			,
    t.reserved_at 	as reserved_at			,
    t.premiere_date as premiere_date		,
    t.premiere_time as premiere_time		,
    array_agg(s.seat_number)::int2[] as reserved_seats
from
    tickets t
    left join seats s on s.ticket_id = t.id
    inner join movies m on m.id = t.movie_id
where 
    t.id = $1
    and t.user_id = $2
group by
    t.id, m.id;

-- name: SelectUserTickets :many
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
    reserved_at desc;


