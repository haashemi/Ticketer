package sql

import (
	"database/sql"
	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Queryer interface {
	sqlx.Queryer
	Exec(query string, args ...any) (sql.Result, error)
}

func Migrate(connString string) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
