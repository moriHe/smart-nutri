import { ChangeDetectorRef, Component, Inject, signal } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Meals, RecipeWithoutIngredients } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';
import { BehaviorSubject, Subscription } from 'rxjs';
import { MealsService } from 'services/meals.service';

@Component({
  selector: 'app-create-mealplan-dialog',
  templateUrl: './create-mealplan-dialog.component.html',
  styleUrls: ['./create-mealplan-dialog.component.css']
})
export class CreateMealplanDialogComponent {
  searchQuery = ""
  private recipesSubscription!: Subscription
  recipes: RecipeWithoutIngredients[] = []

  selectedRecipeId?: number = undefined
  portions = 1


  ngOnInit(): void {
    this.recipesSubscription = this.recipesService.getRecipes().subscribe((response: RecipeWithoutIngredients[]) => {
      this.recipes = response
    })
  }

  searchQueryRecipes() {
    if (this.searchQuery == "") {
      return this.recipes
    }

    return this.recipes.filter((recipe) => {
      return recipe.name.toLowerCase().includes(this.searchQuery.toLowerCase())
    })
  }

  selectRecipe(id: number) {
    if (this.selectedRecipeId === id) {
      return
    }
    
    const selectedRecipe = this.recipes.find((recipe) => recipe.id === id)
    if (selectedRecipe) {
      this.selectedRecipeId = selectedRecipe.id
      this.portions =selectedRecipe.defaultPortions
    } 
    
    this.cdr.detectChanges()
  }

  getPortionLabel() {
    if (this.portions === 1) {
      return "Portion"
    }
    return "Portionen"
  }

  increment(event: Event) {
    event.stopPropagation();
    this.portions = this.portions + 1
  }

  decrement(event: Event) {
    event.stopPropagation();
    if (this.portions === 1) {
      return
    }
    this.portions = this.portions - 1
  }

  isSelected(id: number) {
    if (this.selectedRecipeId === id) {
      return true
    }
    return false
  }

  closeDialog(): void {
    this.dialogRef.close({
      recipeId: this.selectedRecipeId,
      portions: this.portions
    });
  }

  ngOnDestroy(): void {
    if (this.recipesSubscription) {
      this.recipesSubscription.unsubscribe();
    }
  }

  constructor(
    private recipesService: RecipesService,
    private cdr: ChangeDetectorRef,
    public mealsService: MealsService,
    public dialogRef: MatDialogRef<CreateMealplanDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: {recipeId: number, portions: number, selectedMeal: Meals}
    ) {}
}
