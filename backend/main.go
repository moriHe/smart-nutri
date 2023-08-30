package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/moriHe/smart-nutri/api"
	"github.com/moriHe/smart-nutri/storage/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	store := postgres.NewStorage(os.Getenv("DATABASE_URL"))
	api.StartGinServer(store, os.Getenv("SERVER_URL"))

}
