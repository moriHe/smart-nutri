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

func StartGinServer(store storage.Storage) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.GET("/recipes", server.HandleGetAllRecipes)
	router.GET("/recipes/:id", server.HandleGetRecipeById)
	router.Run("localhost:8080")
	return server

}

func (s *Server) HandleGetAllRecipes(c *gin.Context) {
	err, recipes := s.store.GetAllRecipes()
	if err != nil {
		fmt.Println("handleGetAllRecipes error")
		return
	}

	fmt.Println(recipes)

	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

func (s *Server) HandleGetRecipeById(c *gin.Context) {
	id := c.Param("id")

	err, recipe := s.store.GetRecipeById(id)
	if err != nil {
		fmt.Println("handleGetRecipeById error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})

}
