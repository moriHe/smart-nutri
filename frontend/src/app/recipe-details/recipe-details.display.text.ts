import { Meals } from "api/recipes/recipes.interface";

export const MealDisplay = {
    [Meals.NONE]: "-",
    [Meals.BREAKFAST]: "Frühstück",
    [Meals.LUNCH]: "Mittagessen",
    [Meals.DINNER]: "Abendessen"
}