import { ShallowMealplanItem } from "api/mealplans/mealplans.interface"
import { Markets, Units } from "api/recipes/recipes.interface"

export type RecipeIngredientItem = {
    id: number
    ingredientId: number
    name: string
    amountPerPortion: number
    unit: Units
}

export const shoppingListCategories = ["TODAY", "REWE", "EDEKA", "BIO_COMPANY", "WEEKLY_MARKET", "ALDI", "LIDL", "NONE"] as const;
export type ShoppingListByCategory = {
    [Key in typeof shoppingListCategories[number]]: ShoppingListCommonProps[]
}

export type ShoppingListCommonProps = {
    market: Markets
    isBio: boolean
    ingredientId: number
    ingredientName: string
    unit: Units | "PARTIAL"
    items: ShoppingListItem[]
    isDueToday: boolean
    totalAmount: number | null
}

export type ShoppingListItem = {
    shoppingListId: number
    recipeName: string
    mealplanDate: Date
    mealplanPortions: number
    amountPerPortion: number
    recipeIngredientId: number
    ingredientUnit: Units
}


export type AddToShoppingList = {
    recipeIngredientId: number
    market: Markets
    isBio: boolean
  }
