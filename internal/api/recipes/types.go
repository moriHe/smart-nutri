package recipes

type Id = int8
type Name = string

type Ingredient struct {
	RecipeIngredientId int32 `json:"id"`
	Id                 `json:"ingredientId"`
	Name               `json:"name"`
}

type Ingredients = []Ingredient
type ShallowRecipe struct {
	Id   `json:"id"`
	Name `json:"name"`
}

type FullRecipe struct {
	ShallowRecipe
	Ingredients `json:"ingredients"`
}

type PostRecipePayload struct {
	Name        `json:"name"`
	Ingredients []Id `json:"ingredients"`
}

type PostIngredientsPayload struct {
	Ingredients []Id `json:"ingredients"`
}
