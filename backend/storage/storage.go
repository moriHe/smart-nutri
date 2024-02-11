package storage

import (
	"github.com/moriHe/smart-nutri/types"
)

type Storage interface {
	GetUser(fireUid string) (*types.User, *types.RequestError)
	GetUserFamilys(userId int) (*[]types.UserFamily, *types.RequestError)
	// header instead of body?
	PostUser(fireUid string) (*int, *types.RequestError)
	PatchUser(userId int, newActiveFamilyId int) *types.RequestError
	PostFamily(name string, userId int) *types.RequestError

	GetAllRecipes(user *types.User) (*[]types.RecipeWithoutIngredients, *types.RequestError)
	GetRecipeById(recipeId string, activeFamilyId *int) (*types.FullRecipe, *types.RequestError)
	PostRecipe(familyId *int, payload types.PostRecipe) (*types.Id, *types.RequestError)
	PostRecipeIngredient(recipeId string, payload types.PostRecipeIngredient) (*int, *types.RequestError)
	PatchRecipeName(recipeId string, payload types.PatchRecipeName) *types.RequestError
	DeleteRecipe(recipeId string) *types.RequestError
	DeleteRecipeIngredient(recipeIngredientId string) *types.RequestError

	GetMealPlan(familyId *int, date string, forShoppingListStr string) (*[]types.ShallowMealPlanItem, *types.RequestError)
	GetMealPlanItem(id string) (*types.FullMealPlanItem, *types.RequestError)
	PostMealPlanItem(familyId *int, payload types.PostMealPlanItem) *types.RequestError
	DeleteMealPlanItem(id string) *types.RequestError

	GetShoppingListSorted(familyId *int) (*types.ShoppingListByategory, *types.RequestError)

	GetMealPlanItemsShoppingList(familyId *int) (*[]types.ShoppingListMealplanItem, *types.RequestError)
	PostShoppingList(payload []types.PostShoppingListMealplanItem, activeFamilyId *int, mealplanId string) *types.RequestError
	DeleteMealPlanItemShoppingList(id string) *types.RequestError
	DeleteShoppingListItems(ids string, familyId *int) *types.RequestError

	GetInvitationLink(user *types.User) (string, *types.RequestError)
	AcceptInvitation(userId int, token string) *types.RequestError
}
