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

func (s *PostgresStorage) GetAllRecipes() (error, *[]types.ShallowRecipe) {
	rows, err := s.db.Query(context.Background(), "select * from recipes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getAllRecipes: %v\n", err)
		return errors.New("AllRecipeErr"), nil
	}

	defer rows.Close()

	var recipes []types.ShallowRecipe

	for rows.Next() {
		var recipe types.ShallowRecipe

		err = rows.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return errors.New("AllRecipeErr"), nil
		}
		recipes = append(recipes, recipe)
	}

	return nil, &recipes
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

func (s *PostgresStorage) PostRecipe(payload types.PostRecipePayload) error {
	var recipeId int32
	err := s.db.QueryRow(context.Background(), "insert into recipes (name) values ($1) returning id", payload.Name).Scan(&recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert recipe row: %v\n", err)
		return errors.New("post recipe error")
	}

	for i := 0; i < len(payload.Ingredients); i++ {
		_, err := s.db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.Ingredients[i])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
			return errors.New("post recipe loop error")
		}
	}
	return nil
}

func (s *PostgresStorage) PostRecipeIngredient(recipeId string, payload types.PostIngredientsPayload) error {
	for i := 0; i < len(payload.Ingredients); i++ {
		fmt.Println(recipeId, payload.Ingredients[i])
		_, err := s.db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.Ingredients[i])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
			return errors.New("postIngredient error")
		}
	}
	return nil
}

func (s *PostgresStorage) PatchRecipe(recipeId string, payload types.PostRecipePayload) error {
	_, err := s.db.Exec(context.Background(), "update recipes set name = $1 where id = $2", payload.Name, recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to update recipe row: %v\n", err)
		return errors.New("PatchRecipe error")
	}

	return nil
}

func (s *PostgresStorage) DeleteRecipe(recipeId string) error {
	_, err1 := s.db.Exec(context.Background(), "delete from recipes_ingredients where recipe_id =$1", recipeId)

	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err1)
		return errors.New("deleteRecipe error")
	}

	_, err := s.db.Exec(context.Background(), "delete from recipes where id = $1", recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return errors.New("deleteRecipe error")

	}
	return nil
}
