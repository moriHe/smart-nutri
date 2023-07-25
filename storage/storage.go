package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetAllRecipes() (error, *[]types.ShallowRecipe)
	GetRecipeById(string) (error, *types.FullRecipe)
	PostRecipe(types.PostRecipe) error
	PostRecipeIngredient(string, types.PostRecipeIngredient) error
	PatchRecipeName(string, types.PatchRecipeName) error
	DeleteRecipe(string) error
	DeleteRecipeIngredient(string) error
}
