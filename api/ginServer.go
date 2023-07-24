package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
)

type Server struct {
	store storage.Storage
}

func StartServer(store storage.Storage) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.GET("/recipes/:id", server.handleGetUserById)
	router.Run("localhost:8080")
	return server

}

func (s *Server) handleGetUserById(c *gin.Context) {
	id := c.Param("id")

	err, recipe := s.store.GetRecipeById(id)
	if err != nil {
		fmt.Println("Test3")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})

}
