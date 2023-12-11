export interface ShallowRecipe {
    id: number,
    name: string
  }



export enum Meals {
  'NONE' = 'NONE',
  'BREAKFAST' = 'BREAKFAST',
  'LUNCH' = 'LUNCH',
  'DINNER' = 'DINNER'
}

type Markets = 'NONE' |'REWE' | 'EDEKA' | 'BIO_COMPANY' | 'WEEKLY_MARKET' | 'ALDI' | 'LIDL'
type Units = 'GRAM' | 'MILLILITER' | 'TABLESPOON' | 'TEASPOON'

interface RecipeIngredients {
  id: number,
  amountPerPortion: number,
  isBio: boolean,
  market: Markets,
  name: string,
  unit: Units
}
export interface FullRecipe extends ShallowRecipe {
  defaultMeal: Meals,
  defaultPortions: number,
  recipeIngredients: RecipeIngredients
}