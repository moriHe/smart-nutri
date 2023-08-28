package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealPlanRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/mealPlan/:date", s.handleGetMealPlan)
	r.GET("/mealPlan/item/:id", s.handleGetMealPlanItem)
	r.POST("/familys/:familyId/mealPlan", s.handlePostMealPlanItem)
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

func (s *Server) handlePostMealPlanItem(c *gin.Context) {
	familyId := c.Param("familyId")
	var payload types.PostMealPlanItem

	if err := c.BindJSON(&payload); err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		handleResponse[string](c, "Added recipe", s.store.PostMealPlanItem(familyId, payload))
	}
}
