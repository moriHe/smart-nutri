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

export enum Markets {
  'NONE' = 'NONE',
  'REWE' = 'REWE',
  'EDEKA' = 'EDEKA',
  'BIO_COMPANY' = 'BIO_COMPANY',
  'WEEKLY_MARKET' = 'WEEKLY_MARKET',
  'ALDI' = 'ALDI',
  'LIDL' = 'LIDL'
}


export enum Units {
  'GRAM' = 'GRAM',
  'MILLILITER' = 'MILLILITER',
  'TABLESPOON' = 'TABLESPOON',
  'TEASPOON' = 'TEASPOON'
}

export interface RecipeIngredient {
  id: number,
  amountPerPortion: number,
  isBio: boolean,
  market: Markets,
  name: string,
  unit: Units
}

export type RecipeIngredientBody = Omit<RecipeIngredient, "name" | "id"> & {
  ingredientId: number
}

export interface FullRecipe extends ShallowRecipe {
  defaultMeal: Meals,
  defaultPortions: number,
  recipeIngredients: RecipeIngredient[]
}

export type RecipeBody = Omit<FullRecipe, "id">