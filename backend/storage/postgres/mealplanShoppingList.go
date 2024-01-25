package postgres

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/moriHe/smart-nutri/types"
)

var newQuery = "select shopping_list.id, markets.name, shopping_list.is_bio, recipes.name, mealplans.date, mealplans.portions, mealplans.is_shopping_list_item, recipes_ingredients.id, " +
	"ingredients.id, ingredients.name, recipes_ingredients.amount_per_portion, units.name from shopping_list " +
	"left join mealplans on mealplan_id = mealplans.id left join recipes on mealplans.recipe_id = recipes.id left join recipes_ingredients on " +
	"recipes_ingredients_id = recipes_ingredients.id left join meals on mealplans.meal = meals.id left join units on recipes_ingredients.unit = units.id " +
	"left join markets on shopping_list.market = markets.id left join ingredients on recipes_ingredients.ingredient_id = ingredients.id " +
	"where shopping_list.family_id = $1;"

	// type ShoppingList struct {

	// }

func (s *Storage) GetShoppingListSorted(familyId *int) (*types.ShoppingListByategory, error) {
	currentDate := time.Now().UTC()

	rows, _ := s.Db.Query(context.Background(), newQuery, familyId)
	defer rows.Close()
	shoppingList := []types.ShoppingListItemsCommonProps{}

	for rows.Next() {
		var item types.ScanShoppingList
		err := rows.Scan(
			&item.ShoppingListId,
			&item.Market,
			&item.IsBio,
			&item.RecipeName,
			&item.MealplanDate,
			&item.MealPlanPortions,
			&item.IsShoppingListItem,
			&item.RecipeIngredientId,
			&item.IngredientId,
			&item.IngredientName,
			&item.IngredientAmountPerPortion,
			&item.IngredientUnit,
		)
		if err != nil {
			return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Scan Get menu plan shopping list failed: %s", err)}
		}

		roundedAmount := math.Round(float64(item.IngredientAmountPerPortion)*float64(item.MealPlanPortions)*10) / 10
		isDueToday := !item.MealplanDate.IsZero() &&
			item.MealplanDate.Year() == currentDate.Year() &&
			item.MealplanDate.Month() == currentDate.Month() &&
			item.MealplanDate.Day() == currentDate.Day()

		found := false
		for i, existingItem := range shoppingList {
			if existingItem.Market == item.Market &&
				existingItem.IsBio == item.IsBio &&
				existingItem.IngredientId == item.IngredientId &&
				existingItem.IngredientUnit == item.IngredientUnit {
				// Matching item found, append to its Items array
				amount := math.Round((float64(*shoppingList[i].TotalAmount)+roundedAmount)*10) / 10
				shoppingList[i].IsDueToday = shoppingList[i].IsDueToday || isDueToday
				shoppingList[i].TotalAmount = &amount
				shoppingList[i].ShoppingListIds = append(shoppingList[i].ShoppingListIds, item.ShoppingListId)
				shoppingList[i].Items = append(existingItem.Items, types.ShoppingListItem{
					ShoppingListId:                   item.ShoppingListId,
					RecipeName:                       item.RecipeName,
					MealplanDate:                     item.MealplanDate,
					MealPlanPortions:                 item.MealPlanPortions,
					RecipeIngredientAmountPerPortion: item.IngredientAmountPerPortion,
					RecipeIngredientId:               item.RecipeIngredientId,
					RecipeIngredientUnit:             item.IngredientUnit,
				})
				found = true
				break
			} else if existingItem.Market == item.Market &&
				existingItem.IsBio == item.IsBio &&
				existingItem.IngredientId == item.IngredientId &&
				existingItem.IngredientUnit == "PARTIAL" &&
				(item.IngredientUnit == "TABLESPOON" || item.IngredientUnit == "TEASPOON") {
				shoppingList[i].IsDueToday = shoppingList[i].IsDueToday || isDueToday
				shoppingList[i].ShoppingListIds = append(shoppingList[i].ShoppingListIds, item.ShoppingListId)
				shoppingList[i].Items = append(existingItem.Items, types.ShoppingListItem{
					ShoppingListId:                   item.ShoppingListId,
					RecipeName:                       item.RecipeName,
					MealplanDate:                     item.MealplanDate,
					MealPlanPortions:                 item.MealPlanPortions,
					RecipeIngredientAmountPerPortion: item.IngredientAmountPerPortion,
					RecipeIngredientId:               item.RecipeIngredientId,
					RecipeIngredientUnit:             item.IngredientUnit,
				})
				found = true
				break
			}
		}

		if !found && (item.IngredientUnit == "TABLESPOON" || item.IngredientUnit == "TEASPOON") {
			shoppingList = append(shoppingList, types.ShoppingListItemsCommonProps{
				ShoppingListIds: []int{item.ShoppingListId},
				Market:          item.Market,
				IsBio:           item.IsBio,
				IngredientId:    item.IngredientId,
				IngredientName:  item.IngredientName,
				IngredientUnit:  "PARTIAL",
				Items: []types.ShoppingListItem{
					{
						ShoppingListId:                   item.ShoppingListId,
						RecipeName:                       item.RecipeName,
						MealplanDate:                     item.MealplanDate,
						MealPlanPortions:                 item.MealPlanPortions,
						RecipeIngredientAmountPerPortion: item.IngredientAmountPerPortion,
						RecipeIngredientId:               item.RecipeIngredientId,
						RecipeIngredientUnit:             item.IngredientUnit,
					},
				},
				IsDueToday:  isDueToday,
				TotalAmount: nil,
			})
		} else if !found {
			shoppingList = append(shoppingList, types.ShoppingListItemsCommonProps{
				ShoppingListIds: []int{item.ShoppingListId},
				Market:          item.Market,
				IsBio:           item.IsBio,
				IngredientId:    item.IngredientId,
				IngredientName:  item.IngredientName,
				IngredientUnit:  item.IngredientUnit,
				Items: []types.ShoppingListItem{
					{
						ShoppingListId:                   item.ShoppingListId,
						RecipeName:                       item.RecipeName,
						MealplanDate:                     item.MealplanDate,
						MealPlanPortions:                 item.MealPlanPortions,
						RecipeIngredientAmountPerPortion: item.IngredientAmountPerPortion,
						RecipeIngredientId:               item.RecipeIngredientId,
						RecipeIngredientUnit:             item.IngredientUnit,
					},
				},
				IsDueToday:  isDueToday,
				TotalAmount: &roundedAmount,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Something went wrong: %s", err)}
	}

	categorizedItems := types.ShoppingListByategory{
		TODAY:         []types.ShoppingListItemsCommonProps{},
		REWE:          []types.ShoppingListItemsCommonProps{},
		EDEKA:         []types.ShoppingListItemsCommonProps{},
		BIO_COMPANY:   []types.ShoppingListItemsCommonProps{},
		WEEKLY_MARKET: []types.ShoppingListItemsCommonProps{},
		ALDI:          []types.ShoppingListItemsCommonProps{},
		LIDL:          []types.ShoppingListItemsCommonProps{},
		NONE:          []types.ShoppingListItemsCommonProps{},
	}
	for _, item := range shoppingList {
		if item.IsDueToday {
			categorizedItems.TODAY = append(categorizedItems.TODAY, item)
		} else {
			switch item.Market {
			case "REWE":
				categorizedItems.REWE = append(categorizedItems.REWE, item)
			case "EDEKA":
				categorizedItems.EDEKA = append(categorizedItems.EDEKA, item)
			case "BIO_COMPANY":
				categorizedItems.BIO_COMPANY = append(categorizedItems.BIO_COMPANY, item)
			case "WEEKLY_MARKET":
				categorizedItems.WEEKLY_MARKET = append(categorizedItems.WEEKLY_MARKET, item)
			case "ALDI":
				categorizedItems.ALDI = append(categorizedItems.ALDI, item)
			case "LIDL":
				categorizedItems.LIDL = append(categorizedItems.LIDL, item)
			default:
				categorizedItems.NONE = append(categorizedItems.NONE, item)
			}
		}
	}

	return &categorizedItems, nil
}

var getQuery = "select shopping_list.id, mealplans.is_shopping_list_item, mealplans.id, markets.name, shopping_list.is_bio, recipes.id, recipes.name, mealplans.date, mealplans.portions, meals.meal, recipes_ingredients.id, " +
	"ingredients.name, ingredients.id, recipes_ingredients.amount_per_portion, units.name from shopping_list " +
	"left join mealplans on mealplan_id = mealplans.id left join recipes on mealplans.recipe_id = recipes.id left join recipes_ingredients on " +
	"recipes_ingredients_id = recipes_ingredients.id left join meals on mealplans.meal = meals.id left join units on recipes_ingredients.unit = units.id " +
	"left join markets on shopping_list.market = markets.id left join ingredients on recipes_ingredients.ingredient_id = ingredients.id " +
	"where shopping_list.family_id = $1;"

func (s *Storage) GetMealPlanItemsShoppingList(familyId *int) (*[]types.ShoppingListMealplanItem, error) {

	rows, _ := s.Db.Query(context.Background(), getQuery, familyId)
	defer rows.Close()
	shoppingList := []types.ShoppingListMealplanItem{}

	for rows.Next() {
		var item types.ShoppingListMealplanItem
		err := rows.Scan(
			&item.Id,
			&item.MealplanItem.IsShoppingListItem,
			&item.MealplanItem.Id,
			&item.Market,
			&item.IsBio,
			&item.MealplanItem.RecipeId,
			&item.MealplanItem.RecipeName,
			&item.MealplanItem.Date,
			&item.MealplanItem.Portions,
			&item.MealplanItem.Meal,
			&item.RecipeIngredient.Id,
			&item.RecipeIngredient.Name,
			&item.RecipeIngredient.IngredientId,
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

// maybe of use for posting individual items
// func (s *Storage) PostMealPlanItemShoppingList(payload types.PostShoppingListMealplanItem) error {
// 	var marketId int
// 	err := s.Db.QueryRow(context.Background(), "select (id) from markets where markets.name = $1", payload.Market).Scan(&marketId)
// 	if err != nil {
// 		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Step 1: Failed to find market name: %s", err)}
// 	}
// 	_, err = s.Db.Exec(context.Background(), "insert into shopping_list (family_id, mealplan_id, recipes_ingredients_id, market, is_bio) values ($1, $2, $3, $4, $5)", &payload.FamilyId, &payload.MealplanId, &payload.RecipeIngredientId, &marketId, &payload.IsBio)
// 	if err != nil {
// 		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Error: Failed to post mealplan item shopping list: %s", err)}
// 	}

// 	return nil
// }

func (s *Storage) PostShoppingList(payload []types.PostShoppingListMealplanItem, activeFamilyId *int, mealplanId string) error {
	// Start a database transaction
	tx, err := s.Db.Begin(context.Background())
	if err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Failed to start database transaction: %s", err)}
	}
	defer tx.Rollback(context.Background())

	// Iterate through the items and insert them into the database
	for _, item := range payload {
		var marketID int
		err := s.Db.QueryRow(context.Background(), "SELECT id FROM markets WHERE name = $1", item.Market).Scan(&marketID)
		if err != nil {
			return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Failed to find market name: %s", err)}
		}

		_, err = tx.Exec(context.Background(), "insert into shopping_list (family_id, mealplan_id, recipes_ingredients_id, market, is_bio) VALUES ($1, $2, $3, $4, $5)", activeFamilyId, mealplanId, item.RecipeIngredientId, marketID, item.IsBio)
		if err != nil {
			return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprintf("Error: Failed to post mealplan item shopping list: %s", err)}
		}

		_, err = tx.Exec(context.Background(), "update mealplans set is_shopping_list_item = true where family_id = $1 and mealplans.id = $2", activeFamilyId, mealplanId)
	}

	// Commit the transaction if all insertions are successful
	if err := tx.Commit(context.Background()); err != nil {
		return &types.RequestError{Status: http.StatusInternalServerError, Msg: fmt.Sprintf("Failed to commit transaction: %s", err)}
	}

	return nil
}

// TODO: Portions needs to be in mealplanItem
func (s *Storage) DeleteMealPlanItemShoppingList(id string) error {
	item, err := s.Db.Exec(context.Background(), "delete from shopping_list where shopping_list.id = $1", id)

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to delete shopping list item: %v\n", err))
	}

	if item.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Shopping list item does not exist")}
	}

	return nil
}

func (s *Storage) DeleteShoppingListItems(ids string, familyId *int) error {
	if len(ids) == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("No id or ids provided")}
	}

	query := fmt.Sprintf("delete from shopping_list where shopping_list.id in (%s) and family_id = $1", ids)

	item, err := s.Db.Exec(context.Background(), query, &familyId)
	if err != nil {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Unable to delete shopping list item")}
	}

	if item.RowsAffected() == 0 {
		return &types.RequestError{Status: http.StatusBadRequest, Msg: fmt.Sprint("Shopping list item(s) do not exist")}
	}

	return nil
}
