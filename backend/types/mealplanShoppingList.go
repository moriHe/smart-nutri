package types

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
