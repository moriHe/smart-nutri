package types

type ShallowMealPlanItem struct {
	Id         int     `json:"id"`
	RecipeId   int     `json:"recipeId"`
	RecipeName string  `json:"recipeName"`
	Date       string  `json:"date"`
	Portions   float32 `json:"portions"`
	Meal       string  `json:"meal"`
}

// FullRecipe needs amount needs to be multiplied by portions
type MealPlanItemRecipe struct {
	Recipeid          int                `json:"id"`
	Name              string             `json:"name"`
	RecipeIngredients []RecipeIngredient `json:"recipeIngredients"`
}

type FullMealPlanItem struct {
	Id       string             `json:"id"`
	Date     string             `json:"date"`
	Meal     string             `json:"meal"`
	Portions float32            `json:"portions"`
	Recipe   MealPlanItemRecipe `json:"recipe"`
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
