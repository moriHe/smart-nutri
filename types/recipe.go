package types

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
