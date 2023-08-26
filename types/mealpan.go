package types

// TODO: Mealplan endpoints
type ShallowMealPlanItem struct {
	Id         int    `json:"id"`
	RecipeName string `json:"recipeName"`
	Date       string `json:"date"`
	Meal       string `json:"meal"`
}

// FullRecipe needs amount needs to be multiplied by portions
type MealPlanItemRecipe struct {
	Recipeid          int
	Name              string
	RecipeIngredients []RecipeIngredient
}

type FullMealPlanItem struct {
	Id       string
	Date     string
	Meal     string
	Portions float32
	Recipe   MealPlanItemRecipe
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
