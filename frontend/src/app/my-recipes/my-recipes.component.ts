import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { Router } from '@angular/router';
import { Response } from 'api';
import { Meals, ShallowRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';

@Component({
  selector: 'app-my-recipes',
  templateUrl: './my-recipes.component.html',
  styleUrls: ['./my-recipes.component.css']
})
export class MyRecipesComponent {
  nameOfNewRecipe: string = ""
  recipes: ShallowRecipe[] = []

  newRecipeForm = this.formBuilder.group({
    name: ""
  })

  ngOnInit(): void {
    this.recipesService.getRecipes().subscribe((response: ShallowRecipe[]) => {
      this.recipes = response
    })
  }


  onAddRecipe() {
    if (!this.newRecipeForm.value.name) {
      return
    }

    this.recipesService.addRecipe({
      name: this.newRecipeForm.value.name,
      defaultMeal: Meals.BREAKFAST,
      defaultPortions: 1,
      recipeIngredients: []
    }).subscribe((response: Response<{id: number}>) => {
      this.router.navigateByUrl(`rezept/${response.data.id}`)
    })
  }

  openRecipe(id: number) {
    this.router.navigateByUrl(`rezept/${id}`)
  }

  openAddRecipeModal() {
  }


  constructor(
    private recipesService: RecipesService, 
    private router: Router,
    private formBuilder: FormBuilder
    ) { }

}

