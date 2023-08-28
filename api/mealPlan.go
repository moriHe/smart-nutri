package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealPlanRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/mealplan/:date", s.handleGetMealPlan)
	r.GET("/mealPlan/item/:id", s.handleGetMealPlanItem)
	r.POST("/familys/:familyId/mealplan", s.handlePostMealPlanItem)
	r.DELETE("/mealPlan/item/:id", s.handleDeleteMealPlanItem)
}

func (s *Server) handleGetMealPlan(c *gin.Context) {
	familyId := c.Param("familyId")
	date := c.Param("date")
	_, err := time.Parse("2006-01-02", date)

	if err != nil {
		errorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Invalid Date. Use format YYYY-MM-DD"})
		return
	}

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
		handleResponse[string](c, "Added mealplan item", s.store.PostMealPlanItem(familyId, payload))
	}
}

func (s *Server) handleDeleteMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	handleResponse[string](c, "Deleted mealplan item", s.store.DeleteMealPlanItem(id))
}
