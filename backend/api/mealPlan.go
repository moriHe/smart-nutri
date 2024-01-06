package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealPlanRoutes(r *gin.Engine) {
	r.GET("/mealplan/:date", s.handleGetMealPlan)
	r.GET("/mealplan/item/:id", s.handleGetMealPlanItem)
	r.POST("/mealplan", s.handlePostMealPlanItem)
	r.DELETE("/mealplan/item/:id", s.handleDeleteMealPlanItem)
}

func (s *Server) handleGetMealPlan(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	dateStr := c.Param("date")
	date, err := time.Parse(time.RFC3339, dateStr)
	formattedTimestamp := date.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05.999")

	if err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Invalid UTC Date"})
		return
	}

	mealPlan, err := s.store.GetMealPlan(user.ActiveFamilyId, formattedTimestamp)
	responses.HandleResponse[*[]types.ShallowMealPlanItem](c, mealPlan, err)
}

func (s *Server) handleGetMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	mealPlanItem, err := s.store.GetMealPlanItem(id)

	responses.HandleResponse[*types.FullMealPlanItem](c, mealPlanItem, err)
}

func (s *Server) handlePostMealPlanItem(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	var payload types.PostMealPlanItem

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		responses.HandleResponse[string](c, "Added mealplan item", s.store.PostMealPlanItem(user.ActiveFamilyId, payload))
	}
}

func (s *Server) handleDeleteMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	responses.HandleResponse[string](c, "Deleted mealplan item", s.store.DeleteMealPlanItem(id))
}
