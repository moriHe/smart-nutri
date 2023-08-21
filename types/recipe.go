package types

type RecipeIngredient struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	AmountPerPortion string `json:"amountPerPortion"`
	Unit             string `json:"unit"`
	Market           string `json:"market"`
	IsBio            bool   `json:"isBio"`
}

// Posts new ingredient to existing recipe
type PostRecipeIngredient struct {
	IngredientId     int     `json:"ingredientId"`
	AmountPerPortion float32 `json:"amountPerPortion"`
	Unit             int     `json:"unit"`
	Market           int     `json:"market"`
	IsBio            bool    `json:"isBio"`
}

type ShallowRecipe struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FullRecipe struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	RecipeIngredients []RecipeIngredient `json:"recipeIngredients"`
}

type PostRecipe struct {
	Name          string `json:"name"`
	IngredientIds []int  `json:"ingredientIds"`
}

type PatchRecipeName struct {
	Name string `json:"name"`
}
