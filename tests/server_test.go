package tests

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api"
)

func startServer() *gin.Engine {
	r := api.StartGinServer(Db, os.Getenv("DOCKER_TEST_SERVER_URL"))
	return r
}
