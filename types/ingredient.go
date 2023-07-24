package types

type Ingredient struct {
	RecipeIngredientId int32 `json:"id"`
	Id                 `json:"ingredientId"`
	Name               `json:"name"`
}

type Ingredients = []Ingredient

type PostIngredientsPayload struct {
	Ingredients []Id `json:"ingredients"`
}
