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

func StartGinServer(store storage.Storage) *Server {
	router := gin.Default()
	server := &Server{store: store}

	router.GET("/recipes", server.HandleGetAllRecipes)
	router.GET("/recipes/:id", server.HandleGetRecipeById)
	router.POST("/recipes", server.HandlePostRecipe)
	router.POST("/recipes/:id/ingredients", server.HandlePostRecipeIngredient)
	router.PATCH("/recipes/:id", server.HandlePatchRecipeName)
	router.DELETE("/recipes/:id", server.HandleDeleteRecipe)
	router.DELETE("/recipes/ingredients/:id", server.HandleDeleteRecipeIngredient)
	router.Run("localhost:8080")
	return server

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

func successResponse[T any](c *gin.Context, response T) {
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (s *Server) HandleGetAllRecipes(c *gin.Context) {
	recipes, err := s.store.GetAllRecipes()
	if errorResponse(c, err) == true {
		return
	}
	successResponse[*[]types.ShallowRecipe](c, recipes)
}

func (s *Server) HandleGetRecipeById(c *gin.Context) {
	id := c.Param("id")

	recipe, err := s.store.GetRecipeById(id)

	if errorResponse(c, err) == true {
		return
	}
	successResponse[*types.FullRecipe](c, recipe)

}

func (s *Server) HandlePostRecipe(c *gin.Context) {
	var payload types.PostRecipe
	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PostRecipe(payload)
}

func (s *Server) HandlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")

	var payload types.PostRecipeIngredient

	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PostRecipeIngredient(recipeId, payload)
}

func (s *Server) HandlePatchRecipeName(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PatchRecipeName

	if err := c.BindJSON(&payload); err != nil {
		return
	}
	s.store.PatchRecipeName(recipeId, payload)
}

func (s *Server) HandleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	s.store.DeleteRecipe(recipeId)

}

func (s *Server) HandleDeleteRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	s.store.DeleteRecipeIngredient(recipeId)

}
