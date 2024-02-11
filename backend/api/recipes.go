package api

import (
	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) recipeRoutes(r *gin.Engine) {
	r.GET("/recipes", s.handleGetAllRecipes)
	r.GET("/recipes/:id", s.handleGetRecipeById)
	r.POST("/recipes", s.handlePostRecipe)
	r.POST("/recipes/:id/recipeingredient", s.handlePostRecipeIngredient)
	r.PATCH("/recipes/:id", s.handlePatchRecipeName)
	r.DELETE("/recipes/:id", s.handleDeleteRecipe)
	r.DELETE("/recipes/recipeingredient/:id", s.handleDeleteRecipeIngredient)
}

func (s *Server) handleGetAllRecipes(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)

	if user.ActiveFamilyId == nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, "No family found"))
		return
	}
	recipes, err := s.store.GetAllRecipes(user)
	responses.HandleResponse(c, recipes, err)
}

func (s *Server) handleGetRecipeById(c *gin.Context) {
	id := c.Param("id")
	user := contextmethods.GetUserFromContext(c)

	if user.ActiveFamilyId == nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, "No family found"))
		return
	}

	recipe, err := s.store.GetRecipeById(id, user.ActiveFamilyId)
	responses.HandleResponse(c, recipe, err)

}

func (s *Server) handlePostRecipe(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)

	var payload types.PostRecipe

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, err.Error()))
		return
	} else {
		response, err := s.store.PostRecipe(user.ActiveFamilyId, payload)
		responses.HandleResponse(c, response, err)
	}
}

func (s *Server) handlePostRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PostRecipeIngredient
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, "Payload malformed"))
		return
	} else {
		id, err := s.store.PostRecipeIngredient(recipeId, payload)
		responses.HandleResponse(c, id, err)
	}
}

func (s *Server) handlePatchRecipeName(c *gin.Context) {
	recipeId := c.Param("id")
	var payload types.PatchRecipeName

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, types.NewRequestError(&types.BadRequestError, "Payload malformed"))
		return
	} else {
		responses.HandleResponse(c, "Recipe name updated", s.store.PatchRecipeName(recipeId, payload))
	}
}

func (s *Server) handleDeleteRecipe(c *gin.Context) {
	recipeId := c.Param("id")
	responses.HandleResponse(c, "Recipe deleted", s.store.DeleteRecipe(recipeId))
}

func (s *Server) handleDeleteRecipeIngredient(c *gin.Context) {
	recipeId := c.Param("id")
	responses.HandleResponse(c, "Recipe ingredient deleted", s.store.DeleteRecipeIngredient(recipeId))
}
