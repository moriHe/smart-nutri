package types

type ShoppingListMealplanItem struct {
	Id               int                 `json:"id"`
	MealplanItem     ShallowMealPlanItem `json:"mealplanItem"`
	RecipeIngredient RecipeIngredient    `json:"recipeIngredient"`
}
