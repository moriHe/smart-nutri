package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
)

type Server struct {
	store storage.Storage
}

func StartGinServer(store storage.Storage) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.GET("/recipes", server.HandleGetAllRecipes)
	router.GET("/recipes/:id", server.HandleGetRecipeById)
	router.POST("/recipes", server.HandlePostRecipe)
	router.POST("/recipes/:id/ingredients", server.HandlePostRecipeIngredient)
	router.PATCH("/recipes/:id", server.HandlePatchRecipe)
	router.DELETE("/recipes/:id", server.HandleDeleteRecipe)
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

func (s *Server) HandlePostRecipe(c *gin.Context) {
	var payload types.PostRecipePayload
	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PostRecipe(payload)
}

func (s *Server) HandlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")

	var payload types.PostIngredientsPayload

	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PostRecipeIngredient(recipeId, payload)
}

// Patch should also update ingredients
func (s *Server) HandlePatchRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PostRecipePayload

	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PatchRecipe(recipeId, payload)
}

func (s *Server) HandleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	s.store.DeleteRecipe(recipeId)

}
