import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ShallowRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';

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

  openAddRecipeModal() {
    console.log("opened")
  }

  constructor(private recipesService: RecipesService, private router: Router) { }

}

