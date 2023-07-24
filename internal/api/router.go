package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/internal/api/recipes"
)

func ConnRouter() {
	router := gin.Default()

	recipes.RecipeRoutes(router)

	router.Run("localhost:8080")

}
