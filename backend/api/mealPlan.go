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
	// error handling if user nil?
	dateStr := c.Param("date")
	forShoppingListStr := c.Query("forShoppingList")
	date, err := time.Parse(time.RFC3339, dateStr)
	formattedTimestamp := date.Truncate(24 * time.Hour).Format("2006-01-02 15:04:05.999")

	if err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Invalid UTC Date"})
		return
	}

	mealPlan, getmealplanErr := s.store.GetMealPlan(user.ActiveFamilyId, formattedTimestamp, forShoppingListStr)
	responses.HandleResponse(c, mealPlan, getmealplanErr)
}

func (s *Server) handleGetMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	mealPlanItem, err := s.store.GetMealPlanItem(id)

	responses.HandleResponse(c, mealPlanItem, err)
}

func (s *Server) handlePostMealPlanItem(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	var payload types.PostMealPlanItem

	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
		return
	} else {
		responses.HandleResponse(c, "Added mealplan item", s.store.PostMealPlanItem(user.ActiveFamilyId, payload))
	}
}

func (s *Server) handleDeleteMealPlanItem(c *gin.Context) {
	id := c.Param("id")

	responses.HandleResponse(c, "Deleted mealplan item", s.store.DeleteMealPlanItem(id))
}
