package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetAllRecipes() (error, *[]types.ShallowRecipe)
	GetRecipeById(string) (error, *types.FullRecipe)
	PostRecipe(types.PostRecipePayload) error
}
