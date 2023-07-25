package recipes

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/internal/db"
)

func deleteIngredient(c *gin.Context) {
	// The id is the main key in this case
	id := c.Param("id")
	_, err := db.Db.Exec(context.Background(), "delete from recipes_ingredients where recipes_ingredients.id = $1", id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
	}
}
