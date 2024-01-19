import { Component, Inject } from '@angular/core';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { PostMealplanPayload } from 'api/mealplans/mealplans.interface';
import { Meals, RecipeWithoutIngredients } from 'api/recipes/recipes.interface';
import { MealsService } from 'services/meals.service';

export type CreateMealplanDialogData = RecipeWithoutIngredients & 
{
  addMealplanItem: (payload: Omit<PostMealplanPayload, "date" | "isShoppingListItem">) => void,
  selectedMeal: Meals,
  meal: Meals
}

@Component({
  selector: 'app-create-mealplan-bottomsheet',
  templateUrl: './create-mealplan-bottomsheet.component.html',
  styleUrls: ['./create-mealplan-bottomsheet.component.css']
})
export class CreateMealplanBottomsheetComponent {
  meals: Meals[] = Object.values(Meals)
  portions: number = this.data.defaultPortions
  
  addMealplanItem(): void {
    this.data.addMealplanItem({recipeId: this.data.id, meal: this.data.selectedMeal, portions: this.portions})
    this._bottomSheetRef.dismiss();
  }

  getPortionLabel() {
    if (this.portions === 1) {
      return "Portion"
    }
    return "Portionen"
  }

  increment() {
    this.portions = this.portions + 1
  }

  decrement() {
    if (this.portions === 1) {
      return
    }
    this.portions = this.portions - 1
  }
  constructor(
    private _bottomSheetRef: MatBottomSheetRef<CreateMealplanBottomsheetComponent>,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: CreateMealplanDialogData,
    public mealsService: MealsService
    ) {}
}
