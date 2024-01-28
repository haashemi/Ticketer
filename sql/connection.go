package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	*sqlx.DB
}

func NewConnection(connString string) (*Connection, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Connection{DB: db}, nil
}

func (conn *Connection) Transaction(calls func(tx Queryer) error) error {
	tx, err := conn.Beginx()
	if err != nil {
		return err
	}

	if err = calls(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	return err
}
