import { ChangeDetectorRef, Component } from '@angular/core';
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
  private recipesSubscription!: Subscription
  recipes: ShallowRecipe[] = []
  selectedRecipeId?: number = undefined

  ngOnInit(): void {
    this.recipesSubscription = this.recipesService.getRecipes().subscribe((response: ShallowRecipe[]) => {
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
    this.selectedRecipeId = id
    this.cdr.detectChanges()
  }

  isSelected(id: number) {
    if (this.selectedRecipeId === id) {
      return true
    }
    return false
  }

  ngOnDestroy(): void {
    if (this.recipesSubscription) {
      this.recipesSubscription.unsubscribe();
    }
  }

  constructor(
    private recipesService: RecipesService,
    private cdr: ChangeDetectorRef
    ) {}
}
