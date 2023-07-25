package recipes

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/internal/db"
)

func postRecipe(c *gin.Context) {
	var payload PostRecipePayload
	var recipeId int32

	if err := c.BindJSON(&payload); err != nil {
		return
	}

	err := db.Db.QueryRow(context.Background(), "insert into recipes (name) values ($1) returning id", payload.Name).Scan(&recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert recipe row: %v\n", err)
		return
	}

	for i := 0; i < len(payload.Ingredients); i++ {
		_, err := db.Db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.Ingredients[i])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
		}
	}

}

func postIngredient(c *gin.Context) {
	recipeId := c.Param("id")

	var payload PostIngredientsPayload

	if err := c.BindJSON(&payload); err != nil {
		return
	}

	for i := 0; i < len(payload.Ingredients); i++ {
		fmt.Println(recipeId, payload.Ingredients[i])
		_, err := db.Db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.Ingredients[i])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
		}
	}
}

func patchRecipe(c *gin.Context) {
	id := c.Param("id")
	var payload PostRecipePayload

	if err := c.BindJSON(&payload); err != nil {
		return
	}
	fmt.Printf("%s\n", payload.Name)
	fmt.Printf("%s\n", id)

	_, err := db.Db.Exec(context.Background(), "update recipes set name = $1 where id = $2", payload.Name, id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update recipe row: %v\n", err)
	}

}

func deleteRecipe(c *gin.Context) {
	id := c.Param("id")
	_, err1 := db.Db.Exec(context.Background(), "delete from recipes_ingredients where recipe_id =$1", id)

	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err1)
	}

	_, err := db.Db.Exec(context.Background(), "delete from recipes where id = $1", id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
	}

}

func deleteIngredient(c *gin.Context) {
	// The id is the main key in this case
	id := c.Param("id")
	_, err := db.Db.Exec(context.Background(), "delete from recipes_ingredients where recipes_ingredients.id = $1", id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
	}
}
