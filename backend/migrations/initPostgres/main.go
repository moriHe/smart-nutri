package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/Users/moritzhettich/prv/smart-nutri/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// TODO bind file path to this file instead of make file
	// Call from root with initPostgresDev script
	m, err := migrate.New("file://migrations/initPostgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Migration error: %s", err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
