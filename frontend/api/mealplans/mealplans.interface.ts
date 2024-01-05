import { Meals } from "api/recipes/recipes.interface"

type ShallowMealplanItem = {
    id: number,
    recipeId: number, 
    recipeName: string,
    date: string,
    portions: number,
    meal: Meals
}

export type Mealplan = ShallowMealplanItem[]