package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/moriHe/smart-nutri/types"
)

func (s *Storage) GetMealPlan(familyId string, date string) (*[]types.ShallowMealPlanItem, error) {
	query := "select mealplans.id, recipes.name, cast(date as text), portions, meals.meal from mealplans " +
		"join recipes on mealplans.recipe_id = recipes.id join meals on mealplans.meal = meals.id " +
		"where mealplans.family_id = $1 and mealplans.date = $2"
	rows, _ := s.Db.Query(context.Background(), query, familyId, date)
	defer rows.Close()

	mealPlan := []types.ShallowMealPlanItem{}

	for rows.Next() {
		var mealPlanItem types.ShallowMealPlanItem
		err := rows.Scan(&mealPlanItem.Id, &mealPlanItem.RecipeName, &mealPlanItem.Date, &mealPlanItem.Portions, &mealPlanItem.Meal)

		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan Get mealPlanItems failed: %s", err)}
		}
		mealPlan = append(mealPlan, mealPlanItem)
	}

	return &mealPlan, nil
}

func (s *Storage) GetMealPlanItem(id string) (*types.FullMealPlanItem, error) {
	var test types.FullMealPlanItem

	query := "select mealplans.id, cast(date as text), meals.meal, portions, recipes.id " +
		"as recipeId, recipes.name, jsonb_agg(jsonb_build_object(" +
		"'id', recipes_ingredients.id, 'name', ingredients.name, 'unit', units.name, " +
		"'amountPerPortion', recipes_ingredients.amount_per_portion, 'isBio', recipes_ingredients.is_bio, " +
		"'market', markets.name" +
		")) as recipesIngredients from mealplans join meals on mealplans.meal = meals.id " +
		"join recipes on mealplans.recipe_id = recipes.id left join recipes_ingredients " +
		"on recipes.id = recipes_ingredients.recipe_id left join ingredients on " +
		"recipes_ingredients.ingredient_id = ingredients.id left join units on " +
		"recipes_ingredients.unit = units.id  left join markets on recipes_ingredients.market = markets.id " +
		"where mealplans.id = $1 " +
		"group by mealplans.id, meals.meal, recipes.id;"

	err := s.Db.QueryRow(context.Background(), query, id).Scan(&test.Id, &test.Date, &test.Meal, &test.Portions, &test.Recipe.Recipeid, &test.Recipe.Name, &test.Recipe.RecipeIngredients)

	if err != nil {
		return nil, &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Bad Request: %s", err)}
	}

	if test.Recipe.RecipeIngredients[0].Id == 0 {
		test.Recipe.RecipeIngredients = make([]types.RecipeIngredient, 0)
	}

	return &test, nil
}

func (s *Storage) PostMealPlanItem(familyId string, payload types.PostMealPlanItem) error {
	var mealId int
	err := s.Db.QueryRow(context.Background(), "select (id) from meals where meals.meal = $1", payload.Meal).Scan(&mealId)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 1: Failed to find meal name: %s", err)}
	}

	_, err = s.Db.Exec(context.Background(), "insert into mealplans (family_id, recipe_id, date, meal, portions) values ($1, $2, $3, $4, $5)", &familyId, &payload.RecipeId, &payload.Date, &mealId, &payload.Portions)

	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Step 2: Failed to create mealplan item: %s", err)}
	}

	return nil
}

func (s *Storage) DeleteMealPlanItem(id string) error {
	shoppingListItem, err := s.Db.Exec(context.Background(), "delete from mealplans_shopping_list where mealplan_id = $1", id)

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to delete mealplan item 1: %v\n", err))
	}

	mealplanItem, err := s.Db.Exec(context.Background(), "delete from mealplans where mealplans.id = $1", id)

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to delete mealplan item 2: %v\n", err))
	}

	if shoppingListItem.RowsAffected() == 0 || mealplanItem.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Mealplan item does not exist")}
	}

	return nil
}
