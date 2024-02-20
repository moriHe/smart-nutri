package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetUser(fireUid string) (*types.User, error)
	GetUserFamilys(userId int) (*[]types.UserFamily, error)
	// header instead of body?
	PostUser(fireUid string) (*int, error)
	PatchUser(userId int, newActiveFamilyId int) error
	DeleteUser(userId int) error
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
	DeleteMealPlanItem(id string) error

	GetShoppingListSorted(familyId *int) (*types.ShoppingListByategory, error)

	GetMealPlanItemsShoppingList(familyId *int) (*[]types.ShoppingListMealplanItem, error)
	PostShoppingList(payload []types.PostShoppingListMealplanItem, activeFamilyId *int, mealplanId string) error
	DeleteMealPlanItemShoppingList(id string) error
	DeleteShoppingListItems(ids string, familyId *int) error

	GetInvitationLink(user *types.User) (string, error)
	AcceptInvitation(userId int, token string) error
}
