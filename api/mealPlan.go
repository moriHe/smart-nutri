package api

import "github.com/gin-gonic/gin"

func (s *Server) mealPlanRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/mealPlan", s.handleGetAllRecipes)
	r.GET("/mealPlan/item/:id", s.handleGetAllRecipes)
	r.POST("/familys/:familyId/mealPlan", s.handlePostRecipe)
	r.PATCH("/mealPlan/item/:id", s.handlePatchRecipeName)
	r.DELETE("/mealPlan/item/:id", s.handleDeleteRecipe)
}
