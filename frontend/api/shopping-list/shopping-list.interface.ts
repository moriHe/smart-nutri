// type ShoppingListMealplanItem struct {
// 	Id               int                          `json:"id"`
// 	Market           string                       `json:"market"`
// 	IsBio            bool                         `json:"isBio"`
// 	MealplanItem     ShallowMealPlanItem          `json:"mealplanItem"`
// 	RecipeIngredient RecipeIngredientShoppingList `json:"recipeIngredient"`
// }

// type RecipeIngredientShoppingList struct {
// 	Id               int     `json:"id"`
// 	Name             string  `json:"name"`
// 	AmountPerPortion float32 `json:"amountPerPortion"`
// 	Unit             string  `json:"unit"`
// }

import { ShallowMealplanItem } from "api/mealplans/mealplans.interface"
import { Markets, Units } from "api/recipes/recipes.interface"

export type RecipeIngredientItem = {
    id: number
    ingredientId: number
    name: string
    amountPerPortion: number
    unit: Units
}

export type ShoppingListItem = {
    id: number
    market: Markets
    isBio: boolean
    mealplanItem: ShallowMealplanItem
    recipeIngredient: RecipeIngredientItem
}

export type ShoppingListItems = ShoppingListItem[]

export type AddToShoppingList = {
    recipeIngredientId: number
    market: Markets
    isBio: boolean
  }
