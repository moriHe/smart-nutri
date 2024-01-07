import { Component } from '@angular/core';
import { ShallowRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-create-mealplan-dialog',
  templateUrl: './create-mealplan-dialog.component.html',
  styleUrls: ['./create-mealplan-dialog.component.css']
})
export class CreateMealplanDialogComponent {
  searchQuery = ""
  portions!: number
  recipes: ShallowRecipe[] = []
  private recipesSubscription!: Subscription

  ngOnInit(): void {
    this.recipesSubscription = this.recipesService.getRecipes().subscribe((response: ShallowRecipe[]) => {
      this.recipes = response
    })
  }


  ngOnDestroy(): void {
    if (this.recipesSubscription) {
      this.recipesSubscription.unsubscribe();
    }
  }

  constructor(private recipesService: RecipesService) {}
}
