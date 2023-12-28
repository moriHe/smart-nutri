package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	PostUser(fireUid types.PostUser) error // should also return user id

	GetAllRecipes(familyId string) (*[]types.ShallowRecipe, error)
	GetRecipeById(recipeId string) (*types.FullRecipe, error)
	PostRecipe(familyId string, payload types.PostRecipe) (*types.Id, error)
	PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) (*int, error)
	PatchRecipeName(recipeId string, payload types.PatchRecipeName) error
	DeleteRecipe(recipeId string) error
	DeleteRecipeIngredient(recipeIngredientId string) error

	GetMealPlan(familyId string, date string) (*[]types.ShallowMealPlanItem, error)
	GetMealPlanItem(id string) (*types.FullMealPlanItem, error)
	PostMealPlanItem(familyId string, payload types.PostMealPlanItem) error
	DeleteMealPlanItem(id string) error

	GetMealPlanItemsShoppingList(familyId string) (*[]types.ShoppingListMealplanItem, error)
	PostMealPlanItemShoppingList(payload types.PostShoppingListMealplanItem) error
	DeleteMealPlanItemShoppingList(id string) error
}
