package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetAllRecipes(string) (*[]types.ShallowRecipe, error)
	GetRecipeById(string) (*types.FullRecipe, error)
	PostRecipe(string, types.PostRecipe) error
	PostRecipeIngredient(string, types.PostRecipeIngredient) error
	PatchRecipeName(string, types.PatchRecipeName) error
	DeleteRecipe(string) error
	DeleteRecipeIngredient(string) error
}
