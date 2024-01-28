package sql

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	FullName string `db:"full_name"`
	Password string `db:"password"`
	IsAdmin  bool   `db:"is_admin"`
}

func InsertUser(db Queryer, name, email, password string) (id int64, err error) {
	err = sqlx.Get(db, &id, `insert into users (full_name, email, password) values($1, $2, $3) returning id`, name, email, password)
	return
}

func SelectUser(db Queryer, id int64) (data User, err error) {
	err = sqlx.Get(db, &data, `select id, full_name, email, password, is_admin from users where id = $1`, id)
	return
}

func SelectUserByEmail(db Queryer, email string) (data User, err error) {
	err = sqlx.Get(db, &data, `select id, full_name, email, password, is_admin from users where email = $1`, email)
	return
}
