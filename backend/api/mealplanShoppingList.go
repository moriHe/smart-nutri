package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealplanShoppingListRoutes(r *gin.Engine) {
	r.GET("/familys/:familyId/mealplan/shopping-list", s.handleGetMealplanItemsShoppingList)
	r.POST("/familys/:familyId/mealplan/:mealplanId/shopping-list", s.handlePostMealPlanItemShoppingList)
	r.DELETE("/mealplan/shopping-list/:id", s.handleDeleteMealPlanItemShoppingList)
}

func (s *Server) handleGetMealplanItemsShoppingList(c *gin.Context) {
	familyId := c.Param("familyId")
	mealplanItemsShoppingList, err := s.store.GetMealPlanItemsShoppingList(familyId)

	responses.HandleResponse[*[]types.ShoppingListMealplanItem](c, mealplanItemsShoppingList, err)
}

func (s *Server) handlePostMealPlanItemShoppingList(c *gin.Context) {
	payload := types.PostShoppingListMealplanItem{FamilyId: c.Param("familyId"), MealplanId: c.Param("mealplanId")}
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
	} else {
		responses.HandleResponse[string](c, "Added mealplan item to shopping list", s.store.PostMealPlanItemShoppingList(payload))
	}
}

func (s *Server) handleDeleteMealPlanItemShoppingList(c *gin.Context) {
	id := c.Param("id")
	responses.HandleResponse[string](c, "Deleted shopping list item", s.store.DeleteMealPlanItemShoppingList(id))
}
