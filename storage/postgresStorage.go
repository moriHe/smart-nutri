package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

func (s *PostgresStorage) GetAllRecipes() (*[]types.ShallowRecipe, error) {
	rows, err := s.db.Query(context.Background(), "select * from recipes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getAllRecipes: %v\n", err)
		return nil, errors.New("AllRecipeErr")
	}

	defer rows.Close()

	var recipes []types.ShallowRecipe

	for rows.Next() {
		var recipe types.ShallowRecipe

		err = rows.Scan(&recipe.Id, &recipe.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan all recipes failed: %v\n", err)
			return nil, errors.New("AllRecipeErr")
		}
		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}

func (s *PostgresStorage) GetRecipeById(id string) (*types.FullRecipe, error) {

	recipe := types.FullRecipe{Ingredients: []types.RecipeIngredient{}}

	err := s.db.QueryRow(context.Background(), "select id, name from recipes where id=$1", id).Scan(&recipe.Id, &recipe.Name)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Bad Request: No recipe found with id %s", id)}
	}

	rows, _ := s.db.Query(context.Background(), "select recipes_ingredients.id, recipes_ingredients.ingredient_id, name from recipes_ingredients join ingredients on recipes_ingredients.ingredient_id = ingredients.id where recipes_ingredients.recipe_id = $1", id)
	defer rows.Close()

	for rows.Next() {
		var ingredient types.RecipeIngredient

		err = rows.Scan(&ingredient.Id, &ingredient.IngredientId, &ingredient.Name)

		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan recipe_ingredients table failed: %s", err)}
		}

		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed getRecipeById: %v\n", err)
		return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("TEST TEST TEST TEST")}
	}

	return &recipe, nil
}

func (s *PostgresStorage) PostRecipe(payload types.PostRecipe) error {
	var recipeId int32
	err := s.db.QueryRow(context.Background(), "insert into recipes (name) values ($1) returning id", payload.Name).Scan(&recipeId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert recipe row: %v\n", err)
		return errors.New("post recipe error")
	}

	for i := 0; i < len(payload.IngredientIds); i++ {
		_, err := s.db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.IngredientIds[i])

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
			return errors.New("post recipe loop error")
		}
	}
	return nil
}

func (s *PostgresStorage) PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) error {
	fmt.Println(payload.IngredientId)
	_, err := s.db.Exec(context.Background(), "insert into recipes_ingredients(recipe_id, ingredient_id) values ($1, $2)", recipeId, payload.IngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert to recipes_ingredients: %v\n", err)
		return errors.New("postIngredient error")
	}
	return nil
}

func (s *PostgresStorage) PatchRecipeName(recipeId string, payload types.PatchRecipeName) error {
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

func (s *PostgresStorage) DeleteRecipeIngredient(recipeIngredientId string) error {
	_, err := s.db.Exec(context.Background(), "delete from recipes_ingredients where recipes_ingredients.id = $1", recipeIngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return errors.New("delee recipe ingredient error")
	}
	return nil
}
