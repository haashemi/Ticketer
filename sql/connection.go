package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
	*pgxpool.Pool
}

func NewConnection(connString string) (*Connection, error) {
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Connection{Pool: conn}, nil
}

func (conn *Connection) Transaction(ctx context.Context, calls func(tx Querier) error) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err = calls(tx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}

	err = tx.Commit(ctx)
	return err
}
