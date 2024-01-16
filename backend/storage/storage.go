package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetUser(fireUid string) (*types.User, error)
	// header instead of body?
	PostUser(fireUid string) (*int, error)
	PostFamily(name string, userId int) error

	GetAllRecipes(user *types.User) (*[]types.RecipeWithoutIngredients, error)
	GetRecipeById(recipeId string, activeFamilyId *int) (*types.FullRecipe, error)
	PostRecipe(familyId *int, payload types.PostRecipe) (*types.Id, error)
	PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) (*int, error)
	PatchRecipeName(recipeId string, payload types.PatchRecipeName) error
	DeleteRecipe(recipeId string) error
	DeleteRecipeIngredient(recipeIngredientId string) error

	GetMealPlan(familyId *int, date string, forShoppingListStr string) (*[]types.ShallowMealPlanItem, error)
	GetMealPlanItem(id string) (*types.FullMealPlanItem, error)
	PostMealPlanItem(familyId *int, payload types.PostMealPlanItem) error
	PostShoppingList(payload []types.PostShoppingListMealplanItem) error
	DeleteMealPlanItem(id string) error

	GetMealPlanItemsShoppingList(familyId *int) (*[]types.ShoppingListMealplanItem, error)
	PostMealPlanItemShoppingList(payload types.PostShoppingListMealplanItem) error
	DeleteMealPlanItemShoppingList(id string) error
}
