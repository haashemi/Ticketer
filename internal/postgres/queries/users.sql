-- name: InsertUser :one
insert into users (full_name, email, password) values($1, $2, $3) returning id;

-- name: SelectUser :one
select id, full_name, email, password from users where id = $1;

-- name: SelectUserByEmail :one
select id, full_name, email, password from users where email = $1;
