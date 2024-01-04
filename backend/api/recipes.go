package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) recipeRoutes(r *gin.Engine) {
	r.GET("/recipes", s.handleGetAllRecipes)
	// todo familyId
	r.GET("/recipes/:id", s.handleGetRecipeById)
	r.POST("/recipes", s.handlePostRecipe)
	r.POST("/recipes/:id/recipeingredient", s.handlePostRecipeIngredient)
	r.PATCH("/recipes/:id", s.handlePatchRecipeName)
	r.DELETE("/recipes/:id", s.handleDeleteRecipe)
	r.DELETE("/recipes/recipeingredient/:id", s.handleDeleteRecipeIngredient)
}

func (s *Server) handleGetAllRecipes(c *gin.Context) {
	user, err := contextmethods.GetUserFromContext(c)

	if err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Invalid Date. Use format YYYY-MM-DD"})
		return
	}
	if user.ActiveFamilyId == nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "No Family"})
	}
	recipes, err := s.store.GetAllRecipes(user)
	responses.HandleResponse[*[]types.ShallowRecipe](c, recipes, err)
}

func (s *Server) handleGetRecipeById(c *gin.Context) {
	id := c.Param("id")

	recipe, err := s.store.GetRecipeById(id)
	responses.HandleResponse[*types.FullRecipe](c, recipe, err)

}

func (s *Server) handlePostRecipe(c *gin.Context) {
	// TODO: search all c.Param("familyId") and replace with user.displayFamilyId
	familyId := c.Param("familyId")
	var payload types.PostRecipe

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: err.Error()})
	} else {
		response, err := s.store.PostRecipe(familyId, payload)
		responses.HandleResponse[*types.Id](c, response, err)
	}
}

func (s *Server) handlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PostRecipeIngredient
	// TODO add test because of change  from err return to int, err return
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		id, err := s.store.PostRecipeIngredient(recipeId, payload)
		responses.HandleResponse[*int](c, id, err)
	}
}

func (s *Server) handlePatchRecipeName(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PatchRecipeName

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		responses.HandleResponse[string](c, "Recipe name updated", s.store.PatchRecipeName(recipeId, payload))
	}
}

// TODO: handlePatchRecipeIngredient (amount, unit, market, isBio)
func (s *Server) handleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	responses.HandleResponse[string](c, "Recipe deleted", s.store.DeleteRecipe(recipeId))
}

func (s *Server) handleDeleteRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	responses.HandleResponse[string](c, "Recipe ingredient deleted", s.store.DeleteRecipeIngredient(recipeId))
}
