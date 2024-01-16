package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	contextmethods "github.com/moriHe/smart-nutri/api/contextMethods"
	"github.com/moriHe/smart-nutri/api/responses"
	"github.com/moriHe/smart-nutri/types"
)

func (s *Server) mealplanShoppingListRoutes(r *gin.Engine) {
	r.GET("/mealplan/shopping-list", s.handleGetMealplanItemsShoppingList)
	// r.POST("/mealplan/:mealplanId/shopping-list", s.handlePostMealPlanItemShoppingList)
	r.POST("/shopping-list/:mealplanId", s.handlePostShoppingList)
	r.DELETE("/mealplan/shopping-list/:id", s.handleDeleteMealPlanItemShoppingList)
}

func (s *Server) handleGetMealplanItemsShoppingList(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	mealplanItemsShoppingList, err := s.store.GetMealPlanItemsShoppingList(user.ActiveFamilyId)

	responses.HandleResponse[*[]types.ShoppingListMealplanItem](c, mealplanItemsShoppingList, err)
}

// func (s *Server) handlePostMealPlanItemShoppingList(c *gin.Context) {
// 	user := contextmethods.GetUserFromContext(c)
// 	payload := types.PostShoppingListMealplanItem{FamilyId: user.ActiveFamilyId, MealplanId: c.Param("mealplanId")}
// 	if err := c.BindJSON(&payload); err != nil {
// 		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
// 	} else {
// 		responses.HandleResponse[string](c, "Added mealplan item to shopping list", s.store.PostMealPlanItemShoppingList(payload))
// 	}
// }

func (s *Server) handlePostShoppingList(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	mealplanId := c.Param("mealplanId")

	var payload []types.PostShoppingListMealplanItem
	if err := c.BindJSON(&payload); err != nil {
		responses.ErrorResponse(c, &types.RequestError{Status: http.StatusBadRequest, Msg: "Payload malformed"})
		return
	}

	responses.HandleResponse[string](c, "Added mealplan item to shopping list", s.store.PostShoppingList(payload, user.ActiveFamilyId, mealplanId))
}

func (s *Server) handleDeleteMealPlanItemShoppingList(c *gin.Context) {
	id := c.Param("id")
	responses.HandleResponse[string](c, "Deleted shopping list item", s.store.DeleteMealPlanItemShoppingList(id))
}
