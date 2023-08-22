package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
)

type Server struct {
	store storage.Storage
}

func StartGinServer(store storage.Storage, url string) *gin.Engine {
	router := gin.Default()
	server := &Server{store: store}

	server.recipeRoutes(router)
	router.Run(url)

	return router

}

func errorResponse(c *gin.Context, err error) bool {
	if err != nil {
		if requestErr, ok := err.(*types.RequestError); ok {
			c.JSON(requestErr.Status, gin.H{"error": requestErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return true
	}

	return false
}

func handleResponse[T any](c *gin.Context, successResponse T, err error) {
	if err != nil {
		if requestErr, ok := err.(*types.RequestError); ok {
			c.JSON(requestErr.Status, gin.H{"error": requestErr})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": successResponse})
}
