package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/moriHe/smart-nutri/types"
)

func (s *Storage) GetAllRecipes(user *types.User) (*[]types.RecipeWithoutIngredients, error) {
	rows, _ := s.Db.Query(context.Background(), "select recipes.id, name, default_portions, meal from recipes join meals on recipes.default_meal = meals.id where family_id=$1", user.ActiveFamilyId)
	// error handling um silent errors zu vermeiden (hatte id anstatt recipes nach einbau von join, was n silent error warf)
	defer rows.Close()

	var recipes []types.RecipeWithoutIngredients

	for rows.Next() {
		var recipe types.RecipeWithoutIngredients

		err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.DefaultPortions, &recipe.DefaultMeal)
		if err != nil {
			return nil, &types.InternalServerError
		}
		recipes = append(recipes, recipe)
	}

	return &recipes, nil
}

func (s *Storage) GetRecipeById(id string, activeFamilyId *int) (*types.FullRecipe, error) {

	recipe := types.FullRecipe{RecipeIngredients: []types.RecipeIngredient{}}
	query := "select recipes.id, name, default_portions, meal from recipes join meals on recipes.default_meal = meals.id where recipes.id = $1 and family_id = $2"
	err := s.Db.QueryRow(context.Background(), query, id, activeFamilyId).Scan(&recipe.Id, &recipe.Name, &recipe.DefaultPortions, &recipe.DefaultMeal)

	if err != nil {
		return nil, &types.BadRequestError
	}

	recipeIngredientsQuery := "select recipes_ingredients.id, ingredients.id, ingredients.name, ingredients.url, amount_per_portion, " +
		"units.name, markets.name, is_bio from recipes_ingredients " +
		"join ingredients on recipes_ingredients.ingredient_id = ingredients.id " +
		"join units on recipes_ingredients.unit = units.id " +
		"join markets on recipes_ingredients.market = markets.id " +
		"where recipes_ingredients.recipe_id = $1"
	rows, _ := s.Db.Query(context.Background(), recipeIngredientsQuery, id)
	defer rows.Close()

	for rows.Next() {
		var ingredient types.RecipeIngredient

		err = rows.Scan(&ingredient.Id, &ingredient.IngredientId, &ingredient.Name, &ingredient.SourceUrl, &ingredient.AmountPerPortion, &ingredient.Unit, &ingredient.Market, &ingredient.IsBio)

		if err != nil {
			return nil, &types.InternalServerError
		}

		recipe.RecipeIngredients = append(recipe.RecipeIngredients, ingredient)
	}

	return &recipe, nil
}

var postRecipeIngredientQuery = "insert into recipes_ingredients(recipe_id, " +
	"ingredient_id, amount_per_portion, unit, market, is_bio) values ($1, $2, $3, $4, $5, $6) returning id"

func (s *Storage) PostRecipe(familyId *int, payload types.PostRecipe) (*types.Id, error) {
	var defaultMealId int
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return nil, &types.InternalServerError
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select (id) from meals where meals.meal = $1", payload.DefaultMeal).Scan(&defaultMealId)
	if err != nil {
		return nil, &types.BadRequestError
	}
	var recipeId int
	if payload.Name == "" {
		return nil, &types.BadRequestError
	}
	defaultPortions := payload.DefaultPortions
	if defaultPortions == 0 {
		defaultPortions = 1
	}

	err = tx.QueryRow(context.Background(), "insert into recipes (family_id, name, default_portions, default_meal) values ($1, $2, $3, $4) returning id", familyId, payload.Name, defaultPortions, defaultMealId).Scan(&recipeId)

	if err != nil {
		return nil, &types.BadRequestError
	}

	for i := 0; i < len(payload.RecipeIngredients); i++ {
		recipeIngredient := payload.RecipeIngredients[i]
		var unitId int
		err := tx.QueryRow(context.Background(), "select (id) from units where units.name = $1", recipeIngredient.Unit).Scan(&unitId)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.InternalServerError
		}

		var marketId int
		err = tx.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", recipeIngredient.Market).Scan(&marketId)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.InternalServerError
		}

		_, err = tx.Exec(context.Background(), postRecipeIngredientQuery, recipeId, recipeIngredient.IngredientId, recipeIngredient.AmountPerPortion, unitId, marketId, recipeIngredient.IsBio)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.InternalServerError
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &types.BadRequestError
	}
	return &types.Id{Id: recipeId}, nil
}

func (s *Storage) PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) (*int, error) {
	var unitId int
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return nil, &types.InternalServerError
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), "select (id) from units where units.name = $1", payload.Unit).Scan(&unitId)

	if err != nil {
		return nil, &types.BadRequestError
	}

	var marketId int
	err = tx.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", payload.Market).Scan(&marketId)
	if err != nil {
		return nil, &types.BadRequestError
	}
	var id int
	err = tx.QueryRow(context.Background(), postRecipeIngredientQuery, recipeId, payload.IngredientId, payload.AmountPerPortion, unitId, marketId, payload.IsBio).Scan(&id)

	if err != nil {
		return nil, &types.BadRequestError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &types.BadRequestError
	}

	return &id, nil
}

func (s *Storage) PatchRecipeName(recipeId string, payload types.PatchRecipeName) error {
	if payload.Name == "" {
		return &types.BadRequestError
	}
	recipe, err := s.Db.Exec(context.Background(), "update recipes set name = $1 where id = $2 returning id", payload.Name, recipeId)
	if err != nil {
		return &types.BadRequestError
	}

	if recipe.RowsAffected() == 0 {
		return &types.BadRequestError
	}

	return nil
}

func (s *Storage) DeleteRecipe(recipeId string) error {
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.InternalServerError
	}
	defer tx.Rollback(context.Background())

	// Delete related records in a single transaction
	_, err = tx.Exec(context.Background(), "DELETE FROM shopping_list WHERE mealplan_id IN (SELECT id FROM mealplans WHERE recipe_id = $1)", recipeId)
	if err != nil {
		return &types.InternalServerError
	}

	_, err = tx.Exec(context.Background(), "DELETE FROM mealplans WHERE recipe_id = $1", recipeId)
	if err != nil {
		return &types.InternalServerError
	}

	_, err = tx.Exec(context.Background(), "DELETE FROM recipes_ingredients WHERE recipe_id = $1", recipeId)
	if err != nil {
		return &types.InternalServerError
	}

	res, err := tx.Exec(context.Background(), "DELETE FROM recipes WHERE id = $1", recipeId)
	if err != nil {
		return &types.InternalServerError
	}

	if res.RowsAffected() == 0 {
		return &types.BadRequestError
	}

	// Commit the transaction
	err = tx.Commit(context.Background())
	if err != nil {
		return &types.BadRequestError
	}

	return nil
}

func (s *Storage) DeleteRecipeIngredient(recipeIngredientId string) error {
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.BadRequestError
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "delete from shopping_list where recipes_ingredients_id = $1", recipeIngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return &types.InternalServerError
	}

	recipeIngredient, err := tx.Exec(context.Background(), "delete from recipes_ingredients where recipes_ingredients.id = $1", recipeIngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return &types.InternalServerError
	}

	if recipeIngredient.RowsAffected() == 0 {
		return &types.BadRequestError
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return &types.BadRequestError
	}

	return nil
}
