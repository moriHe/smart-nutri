package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan recipes table failed: %s", err)}
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
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Bad Request: No recipe found with id %s", id)}
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
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan recipe_ingredients table failed: %s", err)}
		}

		recipe.RecipeIngredients = append(recipe.RecipeIngredients, ingredient)
	}

	return &recipe, nil
}

var postRecipeIngredientQuery = "insert into recipes_ingredients(recipe_id, " +
	"ingredient_id, amount_per_portion, unit, market, is_bio) values ($1, $2, $3, $4, $5, $6) returning id"

func (s *Storage) PostRecipe(familyId *int, payload types.PostRecipe) (*types.Id, error) {
	var defaultMealId int
	err := s.Db.QueryRow(context.Background(), "select (id) from meals where meals.meal = $1", payload.DefaultMeal).Scan(&defaultMealId)
	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to find meals name: %s", err)}
	}
	var recipeId int
	if payload.Name == "" {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("No recipe name specified")}
	}
	defaultPortions := payload.DefaultPortions
	if defaultPortions == 0 {
		defaultPortions = 1
	}

	err = s.Db.QueryRow(context.Background(), "insert into recipes (family_id, name, default_portions, default_meal) values ($1, $2, $3, $4) returning id", familyId, payload.Name, defaultPortions, defaultMealId).Scan(&recipeId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1: Failed to create recipe: %s", err)}
	}

	for i := 0; i < len(payload.RecipeIngredients); i++ {
		recipeIngredient := payload.RecipeIngredients[i]
		var unitId int
		err := s.Db.QueryRow(context.Background(), "select (id) from units where units.name = $1", recipeIngredient.Unit).Scan(&unitId)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 2a: Failed to create recipe: %s", err)}
		}

		var marketId int
		err = s.Db.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", recipeIngredient.Market).Scan(&marketId)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 2a: Failed to create recipe: %s", err)}
		}

		_, err = s.Db.Exec(context.Background(), postRecipeIngredientQuery, recipeId, recipeIngredient.IngredientId, recipeIngredient.AmountPerPortion, unitId, marketId, recipeIngredient.IsBio)

		if err != nil {
			s.DeleteRecipe(strconv.Itoa(recipeId))
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 2b: Failed to create recipe: %s", err)}
		}
	}
	return &types.Id{Id: recipeId}, nil
}

func (s *Storage) PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) (*int, error) {
	var unitId int
	err := s.Db.QueryRow(context.Background(), "select (id) from units where units.name = $1", payload.Unit).Scan(&unitId)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1a: Failed to create recipe_ingredient: %s", err)}
	}

	var marketId int
	err = s.Db.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", payload.Market).Scan(&marketId)
	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 1b: Failed to create recipe_ingredient: %s", err)}
	}
	var id int
	err = s.Db.QueryRow(context.Background(), postRecipeIngredientQuery, recipeId, payload.IngredientId, payload.AmountPerPortion, unitId, marketId, payload.IsBio).Scan(&id)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 2: Failed to post recipe_ingredient: %s", err)}
	}

	return &id, nil
}

func (s *Storage) PatchRecipeName(recipeId string, payload types.PatchRecipeName) error {
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

func (s *Storage) DeleteRecipe(recipeId string) error {

	rows, err := s.Db.Query(context.Background(), "select id from mealplans where recipe_id = $1", recipeId)
	if err != nil {
		return errors.New("delete mealplan recipe error 1")
	}
	defer rows.Close()

	for rows.Next() {
		var mealplanId int
		err := rows.Scan(&mealplanId)
		if err != nil {
			return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Delete recipe failed 1: %s", err)}
		}
		_, err = s.Db.Exec(context.Background(), "delete from shopping_list where mealplan_id = $1", mealplanId)
		if err != nil {
			return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Delete recipe failed 2: %s", err)}
		}
	}

	_, err = s.Db.Exec(context.Background(), "delete from mealplans where recipe_id = $1", recipeId)
	if err != nil {
		return errors.New("delete mealplan recipe error 2")

	}
	_, err = s.Db.Exec(context.Background(), "delete from recipes_ingredients where recipe_id = $1", recipeId)
	// rows affected
	if err != nil {
		return errors.New("deleteRecipe error")
	}

	recipe, err := s.Db.Exec(context.Background(), "delete from recipes where id = $1", recipeId)

	if err != nil {
		return errors.New("deleteRecipe error")

	}

	if recipe.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Recipe does not exist")}
	}

	return nil
}

func (s *Storage) DeleteRecipeIngredient(recipeIngredientId string) error {
	_, err := s.Db.Exec(context.Background(), "delete from shopping_list where recipes_ingredients_id = $1", recipeIngredientId)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete recipe row: %v\n", err)
		return errors.New("delete recipe ingredient because of shopping list error")
	}

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
