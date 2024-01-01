import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Response } from 'api';
import { ShallowRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';
import { CreateRecipeDialogComponent } from '../create-recipe-dialog/create-recipe-dialog.component';

@Component({
  selector: 'app-my-recipes',
  templateUrl: './my-recipes.component.html',
  styleUrls: ['./my-recipes.component.css']
})
export class MyRecipesComponent {
  recipes: ShallowRecipe[] = []

  ngOnInit(): void {
    this.recipesService.getRecipes().subscribe((response: ShallowRecipe[]) => {
      this.recipes = response
    })
  }


  openRecipe(id: number) {
    this.router.navigateByUrl(`rezept/${id}`)
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(CreateRecipeDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      // TODO add error if something is missing
      if (result && result.name && result.defaultPortions) {
        this.recipesService.addRecipe({
          name: result.name,
          defaultMeal: result.selectedMeal,
          defaultPortions: Number(result.defaultPortions),
          recipeIngredients: []
        }).subscribe((response: Response<{id: number}> | null) => {
          if (response) {
          this.router.navigateByUrl(`rezept/${response.data.id}`)
        }
        })
      }
    })

  }



  constructor(
    private recipesService: RecipesService, 
    private router: Router,
    public dialog: MatDialog
    ) { }

}

