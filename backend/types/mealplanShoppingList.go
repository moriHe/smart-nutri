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

type ScanShoppingList struct {
	ShoppingListId             int       `json:"shoppingListId"`
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

type ShoppingListItemsCommonProps struct {
	ShoppingListIds []int              `json:"shoppingListIds"`
	Market          string             `json:"market"`
	IsBio           bool               `json:"isBio"`
	IngredientId    int                `json:"ingredientId"`
	IngredientName  string             `json:"ingredientName"`
	IngredientUnit  string             `json:"unit"`
	Items           []ShoppingListItem `json:"items"`
	IsDueToday      bool               `json:"isDueToday"`
	TotalAmount     *float64           `json:"totalAmount"`
}
type ShoppingListItem struct {
	ShoppingListId                   int       `json:"shoppingListId"`
	RecipeName                       string    `json:"recipeName"`
	MealplanDate                     time.Time `json:"mealplanDate"`
	MealPlanPortions                 float32   `json:"mealplanPortions"`
	RecipeIngredientAmountPerPortion float32   `json:"recipeIngredientAmountPerPortion"`
	RecipeIngredientId               int       `json:"recipeIngredientId"`
	RecipeIngredientUnit             string    `json:"recipeIngredientUnit"`
}

type ShoppingListByategory struct {
	TODAY         []ShoppingListItemsCommonProps `json:"TODAY"`
	REWE          []ShoppingListItemsCommonProps `json:"REWE"`
	EDEKA         []ShoppingListItemsCommonProps `json:"EDEKA"`
	BIO_COMPANY   []ShoppingListItemsCommonProps `json:"BIO_COMPANY"`
	WEEKLY_MARKET []ShoppingListItemsCommonProps `json:"WEEKLY_MARKET"`
	ALDI          []ShoppingListItemsCommonProps `json:"ALDI"`
	LIDL          []ShoppingListItemsCommonProps `json:"LIDL"`
	NONE          []ShoppingListItemsCommonProps `json:"NONE"`
}
