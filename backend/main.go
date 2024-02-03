package main

import (
	"os"

	"github.com/moriHe/smart-nutri/api"
	"github.com/moriHe/smart-nutri/storage/postgres"
)

func main() {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }
	store := postgres.NewStorage(os.Getenv("DATABASE_URL"))
	api.StartGinServer(store, os.Getenv("GIN_SERVER_URL"))

}
