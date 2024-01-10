import { Meals } from "api/recipes/recipes.interface"

type ShallowMealplanItem = {
    id: number,
    recipeId: number, 
    recipeName: string,
    date: string,
    portions: number,
    meal: Meals,
    isShoppingListItem: boolean
}

export type Mealplan = ShallowMealplanItem[]

export type PostMealplanPayload = {
    recipeId: number,
    date: string,
    meal: Meals,
    portions: number,
    isShoppingListItem: boolean
}