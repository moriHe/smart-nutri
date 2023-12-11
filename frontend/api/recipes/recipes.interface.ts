export interface ShallowRecipe {
    id: number,
    name: string
  }



type DefaultMeal = 'NONE' | 'BREAKFAST' | 'LUNCH' | 'DINNER'
type Market = 'NONE' |'REWE' | 'EDEKA' | 'BIO_COMPANY' | 'WEEKLY_MARKET' | 'ALDI' | 'LIDL'
type Unit = 'GRAM' | 'MILLILITER' | 'TABLESPOON' | 'TEASPOON'

interface RecipeIngredients {
  id: number,
  amountPerPortion: number,
  isBio: boolean,
  market: Market,
  name: string,
  unit: Unit
}
export interface FullRecipe extends ShallowRecipe {
  defaultMeal: DefaultMeal,
  defaultPortions: number,
  recipeIngredients: RecipeIngredients
}