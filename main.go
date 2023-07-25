package main

import (
	"github.com/moriHe/smart-nutri/api"
	"github.com/moriHe/smart-nutri/storage"
)

func main() {
	store := storage.NewPostgresStorage()
	api.StartGinServer(store)

}
