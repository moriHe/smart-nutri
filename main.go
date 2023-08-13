package main

import (
	"os"

	"github.com/moriHe/smart-nutri/api"
	"github.com/moriHe/smart-nutri/storage"
)

func main() {
	store := storage.NewPostgresStorage(os.Getenv("DATABASE_URL"))
	api.StartGinServer(store, "localhost:8080")

}
