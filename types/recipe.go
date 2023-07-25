package types

type RecipeIngredient struct {
	Id           int32  `json:"id"`
	IngredientId int32  `json:"ingredientId"`
	Name         string `json:"name"`
}

// Posts new ingredient to existing recipe
type PostRecipeIngredient struct {
	IngredientId int32 `json:"ingredientId"`
}

type ShallowRecipe struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type FullRecipe struct {
	Id          int32              `json:"id"`
	Name        string             `json:"name"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type PostRecipe struct {
	Name          string  `json:"name"`
	IngredientIds []int32 `json:"ingredientIds"`
}

type PatchRecipeName struct {
	Name string `json:"name"`
}
