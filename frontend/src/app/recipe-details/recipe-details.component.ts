import { Component } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { FullRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';
import { MarketsService } from 'services/markets.service';
import { MealsService } from 'services/meals.service';
import { UnitsService } from 'services/units.service';

@Component({
  selector: 'app-recipe-details',
  templateUrl: './recipe-details.component.html',
  styleUrls: ['./recipe-details.component.css']
})
export class RecipeDetailsComponent {
  recipe?: FullRecipe

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      const id = params['id']
      this.recipesService.getRecipe(id).subscribe((response: FullRecipe) => {
        this.recipe = response
      });
    })
  }

  openSearch() {
    if (this.recipe) {
      this.router.navigateByUrl(`suche/${this.recipe.id}`)
    }
  }

  deleteIngredient(ingredientId: number, ingredientName: string) {
    this.recipesService.removeRecipeIngredient(ingredientId).subscribe({
      complete: () => {
        console.log(this.recipe)
        if (this.recipe) {
          this.snackbar.open(
            `Gelöscht: ${ingredientName}`,
            undefined,
            {
              horizontalPosition: "start",
              verticalPosition: "bottom",
              duration: 1500
            }
          )
          this.recipesService.getRecipe(this.recipe.id).subscribe((response: FullRecipe) => {
            this.recipe = response
          })
        }
      }
    })
  }
// TODO: Doesnt redirect anymore
  deleteRecipe() {
    if (this.recipe) {
      this.recipesService.removeRecipe(this.recipe.id).subscribe({
        next: () => {
          if (this.recipe) {
          this.snackbar.open(
            `Gelöscht: ${this.recipe!.name}`,
             undefined, 
             {
            horizontalPosition: "start",
            verticalPosition: "bottom",
            duration: 1500
            })
          }
          this.router.navigateByUrl("meine-rezepte")
        }
      })
    }
  }

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private recipesService: RecipesService, 
    public mealsService: MealsService,
    public unitsService: UnitsService,
    public marketsService: MarketsService,
    private snackbar: MatSnackBar,
    ) { }

}
