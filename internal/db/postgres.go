package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func ConnPostgres() {
	var err error
	Db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

}

func ClosePostgres() {
	Db.Close(context.Background())
}
