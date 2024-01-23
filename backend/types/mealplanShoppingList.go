package types

import "time"

type PostShoppingListMealplanItem struct {
	RecipeIngredientId int    `json:"recipeIngredientId"`
	Market             string `json:"market"`
	IsBio              bool   `json:"isBio"`
}

type ShoppingListMealplanItem struct {
	Id               int                          `json:"id"`
	Market           string                       `json:"market"`
	IsBio            bool                         `json:"isBio"`
	MealplanItem     ShallowMealPlanItem          `json:"mealplanItem"`
	RecipeIngredient RecipeIngredientShoppingList `json:"recipeIngredient"`
}

type ShoppingList struct {
	Id                         int       `json:"id"`
	Market                     string    `json:"market"`
	IsBio                      bool      `json:"isBio"`
	RecipeName                 string    `json:"recipeName"`
	MealplanDate               time.Time `json:"mealplanDate"`
	MealPlanPortions           float32   `json:"mealplanPortions"`
	IsShoppingListItem         bool      `json:"isShoppingListItem"`
	RecipeIngredientId         int       `json:"recipeIngredientId"`
	IngredientId               int       `json:"ingredientId"`
	IngredientName             string    `json:"ingredientName"`
	IngredientAmountPerPortion float32   `json:"amountPerPortion"`
	IngredientUnit             string    `json:"unit"`
}

type TestWrapper struct {
	Market         string     `json:"market"`
	IsBio          bool       `json:"isBio"`
	IngredientId   int        `json:"ingredientId"`
	IngredientName string     `json:"ingredientName"`
	IngredientUnit string     `json:"unit"`
	Items          []TestItem `json:"items"`
	IsDueToday     bool       `json:"isDueToday"`
	TotalAmount    float64    `json:"totalAmount"`
}
type TestItem struct {
	Id                         int       `json:"id"`
	RecipeName                 string    `json:"recipeName"`
	MealplanDate               time.Time `json:"mealplanDate"`
	MealPlanPortions           float32   `json:"mealplanPortions"`
	IngredientAmountPerPortion float32   `json:"amountPerPortion"`
	RecipeIngredientId         int       `json:"recipeIngredientId"`
}
