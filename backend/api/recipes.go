package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) recipeRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/recipes", s.handleGetAllRecipes)
	r.GET("/recipes/:id", s.handleGetRecipeById)
	r.POST("/familys/:familyId/recipes", s.handlePostRecipe)
	r.POST("/recipes/:id/recipeingredient", s.handlePostRecipeIngredient)
	r.PATCH("/recipes/:id", s.handlePatchRecipeName)
	r.DELETE("/recipes/:id", s.handleDeleteRecipe)
	r.DELETE("/recipes/recipeingredient/:id", s.handleDeleteRecipeIngredient)
}

func (s *Server) handleGetAllRecipes(c *gin.Context) {
	familyId := c.Param("familyId")
	recipes, err := s.store.GetAllRecipes(familyId)
	handleResponse[*[]types.ShallowRecipe](c, recipes, err)
}

func (s *Server) handleGetRecipeById(c *gin.Context) {
	id := c.Param("id")

	recipe, err := s.store.GetRecipeById(id)
	handleResponse[*types.FullRecipe](c, recipe, err)

}

func (s *Server) handlePostRecipe(c *gin.Context) {
	familyId := c.Param("familyId")
	var payload types.PostRecipe

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
	} else {
		response, err := s.store.PostRecipe(familyId, payload)
		handleResponse[*types.Id](c, response, err)
	}
}

func (s *Server) handlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PostRecipeIngredient

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Added recipe ingredient", s.store.PostRecipeIngredient(recipeId, payload))
	}
}

func (s *Server) handlePatchRecipeName(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PatchRecipeName

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Recipe name updated", s.store.PatchRecipeName(recipeId, payload))
	}
}

// TODO: handlePatchRecipeIngredient (amount, unit, market, isBio)
func (s *Server) handleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	handleResponse[string](c, "Recipe deleted", s.store.DeleteRecipe(recipeId))
}

func (s *Server) handleDeleteRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	handleResponse[string](c, "Recipe ingredient deleted", s.store.DeleteRecipeIngredient(recipeId))
}
