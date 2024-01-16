import { FullRecipe, Meals } from "api/recipes/recipes.interface"

export type ShallowMealplanItem = {
    id: number,
    recipeId: number, 
    recipeName: string,
    date: string,
    portions: number,
    meal: Meals,
    isShoppingListItem: boolean
}

export type FullMealplanItem = Omit<ShallowMealplanItem, "recipeId" | "recipeName"> & {
    recipe: Pick<FullRecipe, "id" | "name" | "recipeIngredients">
}

export type Mealplans = ShallowMealplanItem[]

export type PostMealplanPayload = {
    recipeId: number,
    date: string,
    meal: Meals,
    portions: number,
    isShoppingListItem: boolean
}
