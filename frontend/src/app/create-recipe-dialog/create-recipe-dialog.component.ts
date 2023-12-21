import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Meals } from 'api/recipes/recipes.interface';
import { MealsService } from 'services/meals.service';

@Component({
  selector: 'app-create-recipe-dialog',
  templateUrl: './create-recipe-dialog.component.html',
  styleUrls: ['./create-recipe-dialog.component.css']
})
export class CreateRecipeDialogComponent {
  name!: string
  selectedMeal: Meals = Meals.NONE
  meals: Meals[] = Object.values(Meals)
  defaultPortions: number = 1

  onNoClick(): void {
    this.dialogRef.close();
  }

  constructor(
    public dialogRef: MatDialogRef<CreateRecipeDialogComponent>,
    public mealsService: MealsService
  ) {}

}
