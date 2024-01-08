package types

type Id struct {
	Id int `json:"id"`
}

type RecipeIngredientShoppingList struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	Unit             string  `json:"unit"`
}

type RecipeIngredient struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	Unit             string  `json:"unit"`
	Market           string  `json:"market"`
	IsBio            bool    `json:"isBio"`
}

type RecipeWithoutIngredients struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	DefaultPortions float32 `json:"defaultPortions"`
	DefaultMeal     string  `json:"defaultMeal"`
}

type FullRecipe struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	DefaultPortions   float32            `json:"defaultPortions"`
	DefaultMeal       string             `json:"defaultMeal"`
	RecipeIngredients []RecipeIngredient `json:"recipeIngredients"`
}
type PostRecipeIngredient struct {
	IngredientId     int     `json:"ingredientId"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	Unit             string  `json:"unit"`
	Market           string  `json:"market"`
	IsBio            bool    `json:"isBio"`
}
type PostRecipe struct {
	Name              string                 `json:"name"`
	DefaultPortions   float32                `json:"defaultPortions"`
	DefaultMeal       string                 `json:"defaultMeal"`
	RecipeIngredients []PostRecipeIngredient `json:"recipeIngredients"`
}

type PatchRecipeName struct {
	Name string `json:"name"`
}
