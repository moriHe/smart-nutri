import { Component, Inject } from '@angular/core';
import { MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { PostMealplanPayload } from 'api/mealplans/mealplans.interface';
import { Meals, RecipeWithoutIngredients } from 'api/recipes/recipes.interface';
import { MealsService } from 'services/meals.service';

type BottomSheetData = RecipeWithoutIngredients & 
{addMealplanItem: (payload: Omit<PostMealplanPayload, "date">) => void}

@Component({
  selector: 'app-create-mealplan-bottomsheet',
  templateUrl: './create-mealplan-bottomsheet.component.html',
  styleUrls: ['./create-mealplan-bottomsheet.component.css']
})
export class CreateMealplanBottomsheetComponent {
  selectedMeal: Meals = this.data.defaultMeal
  meals: Meals[] = Object.values(Meals)
  portions: number = this.data.defaultPortions
  
  addMealplanItem(): void {
    this.data.addMealplanItem({recipeId: this.data.id, meal: this.selectedMeal, portions: this.portions})
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
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: BottomSheetData,
    public mealsService: MealsService
    ) {}
}
