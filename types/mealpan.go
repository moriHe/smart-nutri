package types

// TODO: Mealplan endpoints
type ShallowMealPlanItem struct {
	Id         int    `json:"id"`
	RecipeName string `json:"recipeName"`
	Date       string `json:"date"`
	Meal       string `json:"meal"`
}

// FullRecipe needs amount needs to be multiplied by portions
type FullMealPlanItem struct {
	Id       int        `json:"id"`
	Date     string     `json:"date"`
	Recipe   FullRecipe `json:"Recipe"`
	Meal     string     `json:"meal"`
	Portions float32    `json:"portions"`
}

type PostMealPlanItem struct {
	RecipeId int     `json:"recipeId"`
	Date     string  `json:"date"`
	Meal     string  `json:"meal"`
	Portions float32 `json:"portions"`
}

type PatchMealPlanItem struct {
	Date     string  `json:"date"`
	Meal     string  `json:"meal"`
	Portions float32 `json:"portions"`
}
