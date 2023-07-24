package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetRecipeById(string) (error, *types.FullRecipe)
}
