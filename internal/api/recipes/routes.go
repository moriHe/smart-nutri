package recipes

import (
	"github.com/gin-gonic/gin"
)

func RecipeRoutes(router *gin.Engine) {
	router.GET("/recipes", getAllRecipes)
	router.GET("/recipes/:id", getRecipeById)
	router.POST("/recipes", postRecipe)
	router.PATCH("/recipes/:id", patchRecipe)
	router.DELETE("/recipes/:id", deleteRecipe)

	router.POST("recipes/:id/ingredients", postIngredient)
	router.DELETE("recipes/ingredients/:id", deleteIngredient)
}
