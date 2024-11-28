package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dsnrw, dsnro string) (*sql.DB, *sql.DB) {
	dbrw, err := sql.Open("pgx", dsnrw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to write database: %v\n", err)
		os.Exit(1)
	}

	if err = dbrw.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	dbro, err := sql.Open("pgx", dsnro)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to read-only database: %v\n", err)
		os.Exit(1)
	}

	if err = dbro.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	return dbrw, dbro
}
