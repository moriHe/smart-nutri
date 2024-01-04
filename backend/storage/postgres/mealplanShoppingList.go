package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/moriHe/smart-nutri/types"
)

var getQuery = "select mealplans_shopping_list.id, mealplans.id, markets.name, mealplans_shopping_list.is_bio, recipes.name, cast(mealplans.date as text), mealplans.portions, meals.meal, recipes_ingredients.id, " +
	"ingredients.name, recipes_ingredients.amount_per_portion, units.name from mealplans_shopping_list " +
	"left join mealplans on mealplan_id = mealplans.id left join recipes on mealplans.recipe_id = recipes.id left join recipes_ingredients on " +
	"recipes_ingredients_id = recipes_ingredients.id left join meals on mealplans.meal = meals.id left join units on recipes_ingredients.unit = units.id " +
	"left join markets on mealplans_shopping_list.market = markets.id left join ingredients on recipes_ingredients.ingredient_id = ingredients.id " +
	"where mealplans_shopping_list.family_id = $1;"

func (s *Storage) GetMealPlanItemsShoppingList(familyId *int) (*[]types.ShoppingListMealplanItem, error) {

	rows, _ := s.Db.Query(context.Background(), getQuery, familyId)
	defer rows.Close()
	shoppingList := []types.ShoppingListMealplanItem{}

	for rows.Next() {
		var item types.ShoppingListMealplanItem
		err := rows.Scan(
			&item.Id,
			&item.MealplanItem.Id,
			&item.Market,
			&item.IsBio,
			&item.MealplanItem.RecipeName,
			&item.MealplanItem.Date,
			&item.MealplanItem.Portions,
			&item.MealplanItem.Meal,
			&item.RecipeIngredient.Id,
			&item.RecipeIngredient.Name,
			&item.RecipeIngredient.AmountPerPortion,
			&item.RecipeIngredient.Unit,
		)

		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan Get menu plan shopping list failed: %s", err)}
		}

		shoppingList = append(shoppingList, item)
	}

	return &shoppingList, nil
}

func (s *Storage) PostMealPlanItemShoppingList(payload types.PostShoppingListMealplanItem) error {
	var marketId int
	err := s.Db.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", payload.Market).Scan(&marketId)
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 1: Failed to find market name: %s", err)}
	}
	_, err = s.Db.Exec(context.Background(), "insert into mealplans_shopping_list (family_id, mealplan_id, recipes_ingredients_id, market, is_bio) values ($1, $2, $3, $4, $5)", &payload.FamilyId, &payload.MealplanId, &payload.RecipeIngredientId, &marketId, &payload.IsBio)
	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Error: Failed to post mealplan item shopping list: %s", err)}
	}

	return nil
}

// TODO: Portions needs to be in mealplanItem
func (s *Storage) DeleteMealPlanItemShoppingList(id string) error {
	item, err := s.Db.Exec(context.Background(), "delete from mealplans_shopping_list where mealplans_shopping_list.id = $1", id)

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to delete shopping list item: %v\n", err))
	}

	if item.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Shopping list item does not exist")}
	}

	return nil
}
