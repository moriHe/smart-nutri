package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/moriHe/smart-nutri/types"
)

type PostgresStorage struct {
	db *pgx.Conn
}

func NewPostgresStorage() *PostgresStorage {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) GetRecipeById(id string) (error, *types.FullRecipe) {

	var recipe types.FullRecipe

	err := s.db.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&recipe.Id, &recipe.Name)

	rows, recipeIngredientsErr := s.db.Query(context.Background(), "select recipes_ingredients.id, recipes_ingredients.ingredient_id, name from recipes_ingredients join ingredients on recipes_ingredients.ingredient_id = ingredients.id where recipes_ingredients.recipe_id = $1", id)

	if recipeIngredientsErr != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getAllRecipes: %v\n", err)
		return errors.New("Test"), nil
	}

	defer rows.Close()

	for rows.Next() {
		var ingredient types.Ingredient

		err = rows.Scan(&ingredient.RecipeIngredientId, &ingredient.Id, &ingredient.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return errors.New("Test1"), nil
		}

		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getRecipeById: %v\n", err)
		return errors.New("Test2"), nil
	}

	return nil, &recipe
}
