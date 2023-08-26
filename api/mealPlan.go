package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealPlanRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/mealPlan/:date", s.handleGetMealPlan)
	r.GET("/mealPlan/item/:id", s.handleGetMealPlanItem)
	r.POST("/familys/:familyId/mealPlan", s.handlePostRecipe)
	r.PATCH("/mealPlan/item/:id", s.handlePatchRecipeName)
	r.DELETE("/mealPlan/item/:id", s.handleDeleteRecipe)
}

func (s *Server) handleGetMealPlan(c *gin.Context) {
	familyId := c.Param("familyId")
	date := c.Param("date")

	mealPlan, err := s.store.GetMealPlan(familyId, date)
	handleResponse[*[]types.ShallowMealPlanItem](c, mealPlan, err)
}

func (s *Server) handleGetMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	mealPlanItem, err := s.store.GetMealPlanItem(id)

	handleResponse[*types.FullMealPlanItem](c, mealPlanItem, err)
}
