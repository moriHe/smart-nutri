package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moriHe/smart-nutri/types"
)

type PostgresStorage struct {
	Db *pgxpool.Pool
}

func NewPostgresStorage(url string) *PostgresStorage {
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &PostgresStorage{Db: db}
}

func (s *PostgresStorage) GetAllRecipes() (*[]types.ShallowRecipe, error) {
	rows, _ := s.Db.Query(context.Background(), "select * from recipes")

	defer rows.Close()

	var recipes []types.ShallowRecipe

	for rows.Next() {
		var recipe types.ShallowRecipe

		err := rows.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan recipes table failed: %s", err)}
		}
		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}

func (s *PostgresStorage) GetRecipeById(id string) (*types.FullRecipe, error) {

	recipe := types.FullRecipe{Ingredients: []types.RecipeIngredient{}}

	err := s.Db.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&recipe.Id, &recipe.Name)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Bad Request: No recipe found with id %s", id)}
	}

	recipeIngredientsQuery := "select recipes_ingredients.id, recipes_ingredients.ingredient_id, name from recipes_ingredients join ingredients on recipes_ingredients.ingredient_id = ingredients.id where recipes_ingredients.recipe_id = $1"
	rows, _ := s.Db.Query(context.Background(), recipeIngredientsQuery, id)
	defer rows.Close()

	for rows.Next() {
		var ingredient types.RecipeIngredient

		err = rows.Scan(&ingredient.Id, &ingredient.IngredientId, &ingredient.Name)

		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan recipe_ingredients table failed: %s", err)}
		}

		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	return &recipe, nil
}

func (s *PostgresStorage) PostRecipe(payload types.PostRecipe) error {
	var recipeId int
	if payload.Name == "" {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("No recipe name specified")}
	}
	err := s.Db.QueryRow(context.Background(), "insert into recipes (name) values ($1) returning id", payload.Name).Scan(&recipeId)

	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to create recipe: %s", err)}
	}

	for i := 0; i < len(payload.IngredientIds); i++ {
		_, err := s.Db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.IngredientIds[i])

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 2: Failed to create recipe: %s", err)}
		}
	}
	return nil
}

func (s *PostgresStorage) PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) error {
	_, err := s.Db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.IngredientId)

	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Failed to post recipe_ingredient: %s", err)}
	}
	return nil
}

func (s *PostgresStorage) PatchRecipeName(recipeId string, payload types.PatchRecipeName) error {
	if payload.Name == "" {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("No recipe name specified")}
	}
	recipe, err := s.Db.Exec(context.Background(), "update recipes set name = $1 where id = $2 returning id", payload.Name, recipeId)
	if err != nil {
		return errors.New("PatchRecipe error")
	}

	if recipe.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Recipe does not exist")}
	}

	return nil
}

func (s *PostgresStorage) DeleteRecipe(recipeId string) error {
	_, err1 := s.Db.Exec(context.Background(), "delete from recipes_ingredients where recipe_id =$1", recipeId)

	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err1)
		return errors.New("deleteRecipe error")
	}

	recipe, err := s.Db.Exec(context.Background(), "delete from recipes where id = $1", recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return errors.New("deleteRecipe error")

	}

	if recipe.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Recipe does not exist")}
	}

	return nil
}

func (s *PostgresStorage) DeleteRecipeIngredient(recipeIngredientId string) error {
	recipeIngredient, err := s.Db.Exec(context.Background(), "delete from recipes_ingredients where recipes_ingredients.id = $1", recipeIngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return errors.New("delee recipe ingredient error")
	}

	if recipeIngredient.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Recipe ingredient does not exist")}
	}

	return nil
}
