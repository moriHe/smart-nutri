import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Meals } from 'api/recipes/recipes.interface';
import { MealsService } from 'services/meals.service';

interface RecipeDialogData {
  name: string,
  selectedMeal: Meals
  meals: Meals[]
  defaultPortions: number
}

@Component({
  selector: 'app-create-recipe-dialog',
  templateUrl: './create-recipe-dialog.component.html',
  styleUrls: ['./create-recipe-dialog.component.css']
})
export class CreateRecipeDialogComponent {
  onNoClick(): void {
    this.dialogRef.close();
  }

  constructor(
    public dialogRef: MatDialogRef<CreateRecipeDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: RecipeDialogData,
    public mealsService: MealsService
  ) {}

}
