package types

type RecipeIngredient struct {
	Id           int    `json:"id"`
	IngredientId int    `json:"ingredientId"`
	Name         string `json:"name"`
}

// Posts new ingredient to existing recipe
type PostRecipeIngredient struct {
	IngredientId int `json:"ingredientId"`
}

type ShallowRecipe struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FullRecipe struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type PostRecipe struct {
	Name          string `json:"name"`
	IngredientIds []int  `json:"ingredientIds"`
}

type PatchRecipeName struct {
	Name string `json:"name"`
}
