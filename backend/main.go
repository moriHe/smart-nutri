package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/moriHe/smart-nutri/api"
	"github.com/moriHe/smart-nutri/storage/postgres"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	m, err := migrate.New("file://migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Migration error: %s", err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}

	store := postgres.NewStorage(os.Getenv("DATABASE_URL"))
	api.StartGinServer(store, os.Getenv("PORT"))

}
