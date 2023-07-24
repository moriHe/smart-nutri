package main

import (
	api "github.com/moriHe/smart-nutri/internal/api"
	"github.com/moriHe/smart-nutri/internal/db"
)

func main() {
	db.ConnPostgres()
	defer db.ClosePostgres()

	api.ConnRouter()

}
