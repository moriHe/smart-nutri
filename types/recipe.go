package types

type RecipeIngredient struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	Unit             string  `json:"unit"`
	Market           string  `json:"market"`
	IsBio            bool    `json:"isBio"`
}

// Posts new ingredient to existing recipe

type ShallowRecipe struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FullRecipe struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	RecipeIngredients []RecipeIngredient `json:"recipeIngredients"`
}
type PostRecipeIngredient struct {
	IngredientId     int     `json:"ingredientId"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	UnitId           int     `json:"unitId"`
	MarketId         int     `json:"marketId"`
	IsBio            bool    `json:"isBio"`
}
type PostRecipe struct {
	Name              string                 `json:"name"`
	RecipeIngredients []PostRecipeIngredient `json:"recipeIngredients"`
}

type PatchRecipeName struct {
	Name string `json:"name"`
}
