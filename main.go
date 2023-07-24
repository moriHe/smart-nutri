package main

import (
	api "github.com/moriHe/smart-nutri/internal/api"
	"github.com/moriHe/smart-nutri/internal/db"
)

func main() {
	db.ConnPostgres()
	defer db.ClosePostgres()

	api.ConnRouter()

	/*
		router.GET("/recipes", getAllRecipes)
			router.GET("/recipes/:id", getRecipeById)
			router.POST("/recipes", postRecipe)
			router.PATCH("/recipes/:id", patchRecipe)
			router.DELETE("/recipes/:id", deleteRecipe)

			router.POST("recipes/:id/ingredients", postIngredient)
			router.DELETE("recipes/ingredients/:id", deleteIngredient)
	*/

}
