package recipes

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moriHe/smart-nutri/internal/db"
)

type Id = int8
type Name = string

type Ingredient struct {
	RecipeIngredientId int32 `json:"id"`
	Id                 `json:"ingredientId"`
	Name               `json:"name"`
}

type Ingredients = []Ingredient
type ShallowRecipe struct {
	Id   `json:"id"`
	Name `json:"name"`
}

type FullRecipe struct {
	ShallowRecipe
	Ingredients `json:"ingredients"`
}

func getAllRecipes(c *gin.Context) {
	rows, err := db.Db.Query(context.Background(), "select * from recipes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getAllRecipes: %v\n", err)
		return
	}

	defer rows.Close()

	var recipes []ShallowRecipe

	for rows.Next() {
		var recipe ShallowRecipe

		err = rows.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return
		}
		recipes = append(recipes, recipe)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
	}
	fmt.Print(recipes)
	c.JSON(http.StatusOK, gin.H{
		"data": recipes,
	})

}

func getRecipeById(c *gin.Context) {
	id := c.Param("id")

	var recipe FullRecipe

	err := db.Db.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&recipe.Id, &recipe.Name)

	rows, recipeIngredientsErr := db.Db.Query(context.Background(), "select recipes_ingredients.id, recipes_ingredients.ingredient_id, name from recipes_ingredients join ingredients on recipes_ingredients.ingredient_id = ingredients.id where recipes_ingredients.recipe_id = $1", id)

	if recipeIngredientsErr != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getAllRecipes: %v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient Ingredient

		err = rows.Scan(&ingredient.RecipeIngredientId, &ingredient.Id, &ingredient.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return
		}

		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getRecipeById: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

type PostRecipePayload struct {
	Name        `json:"name"`
	Ingredients []Id `json:"ingredients"`
}

type PostIngredientsPayload struct {
	Ingredients []Id `json:"ingredients"`
}

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

// TODO Delete RecipeIngredient

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
