package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/storage"
	"github.com/moriHe/smart-nutri/types"
)

type Server struct {
	store storage.Storage
	R     *gin.Engine
}

func StartGinServer(store storage.Storage, url string) *Server {
	router := gin.Default()
	server := &Server{store: store, R: router}

	router.GET("/familys/:familyId/recipes", server.HandleGetAllRecipes)
	router.GET("/recipes/:id", server.HandleGetRecipeById)
	router.POST("/familys/:familyId/recipes", server.HandlePostRecipe)
	router.POST("/recipes/:id/ingredients", server.HandlePostRecipeIngredient)
	router.PATCH("/recipes/:id", server.HandlePatchRecipeName)
	router.DELETE("/recipes/:id", server.HandleDeleteRecipe)
	router.DELETE("/recipes/ingredients/:id", server.HandleDeleteRecipeIngredient)
	router.Run(url)
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

func (s *Server) HandleGetAllRecipes(c *gin.Context) {
	familyId := c.Param("familyId")
	recipes, err := s.store.GetAllRecipes(familyId)
	handleResponse[*[]types.ShallowRecipe](c, recipes, err)
}

func (s *Server) HandleGetRecipeById(c *gin.Context) {
	id := c.Param("id")

	recipe, err := s.store.GetRecipeById(id)
	handleResponse[*types.FullRecipe](c, recipe, err)

}

func (s *Server) HandlePostRecipe(c *gin.Context) {
	familyId := c.Param("familyId")
	var payload types.PostRecipe

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Added recipe", s.store.PostRecipe(familyId, payload))
	}
}

func (s *Server) HandlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PostRecipeIngredient

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Added recipe ingredient", s.store.PostRecipeIngredient(recipeId, payload))
	}
}

func (s *Server) HandlePatchRecipeName(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PatchRecipeName

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Recipe name updated", s.store.PatchRecipeName(recipeId, payload))
	}
}

func (s *Server) HandleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	handleResponse[string](c, "Recipe deleted", s.store.DeleteRecipe(recipeId))
}

func (s *Server) HandleDeleteRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	handleResponse[string](c, "Recipe ingredient deleted", s.store.DeleteRecipeIngredient(recipeId))
}
