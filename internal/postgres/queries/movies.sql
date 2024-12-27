-- name: InsertMovie :exec
insert into movies (name, movie_time, genres, premiere_time, premiere_from_date, premiere_to_date) values ($1, $2, $3, $4, $5, $6);

-- name: SelectMovies :many
select id, name, premiere_time from movies order by premiere_time asc;

-- name: SelectMovie :one
select id, name, movie_time, genres, premiere_from_date, premiere_to_date, premiere_time from movies where id = $1;

-- name: SelectMovieReservedSeats :many
select 
    s.seat_number::int2
from 
    tickets t 
    left join seats s on s.ticket_id = t.id
where
    t.movie_id = $1
    and t.premiere_date = $2;
