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
	r.GET("shopping-list", s.handleGetShoppingList)
	r.POST("/shopping-list/:mealplanId", s.handlePostShoppingList)
	r.DELETE("/mealplan/shopping-list/:id", s.handleDeleteMealPlanItemShoppingList)
}

func (s *Server) handleGetMealplanItemsShoppingList(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)
	mealplanItemsShoppingList, err := s.store.GetMealPlanItemsShoppingList(user.ActiveFamilyId)

	responses.HandleResponse[*[]types.ShoppingListMealplanItem](c, mealplanItemsShoppingList, err)
}

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

func (s *Server) handleGetShoppingList(c *gin.Context) {
	user := contextmethods.GetUserFromContext(c)

	shoppingList, err := s.store.GetShoppingListSorted(user.ActiveFamilyId)
	responses.HandleResponse[*types.ShoppingListByategory](c, shoppingList, err)
}
