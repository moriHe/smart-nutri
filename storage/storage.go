package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetAllRecipes() (error, *[]types.ShallowRecipe)
	GetRecipeById(string) (error, *types.FullRecipe)
	PostRecipe(types.PostRecipePayload) error
	PostRecipeIngredient(string, types.PostIngredientsPayload) error
	// TODO Rename PatchRecipe to PatchRecipeName or something like that
	PatchRecipe(string, types.PostRecipePayload) error
	DeleteRecipe(string) error
}
