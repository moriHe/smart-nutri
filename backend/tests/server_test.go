package tests

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api"
)

func startServer() (*gin.Engine, error) {
	r, err := api.StartGinServer(Db, os.Getenv("DOCKER_TEST_SERVER_URL"))

	if err != nil {
		return nil, err
	}
	return r, nil
}
