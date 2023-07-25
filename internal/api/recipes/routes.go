package recipes

import (
	"github.com/gin-gonic/gin"
)

func RecipeRoutes(router *gin.Engine) {

	router.DELETE("recipes/ingredients/:id", deleteIngredient)
}
