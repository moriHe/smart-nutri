import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Recipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';

@Component({
  selector: 'app-my-recipes',
  templateUrl: './my-recipes.component.html',
  styleUrls: ['./my-recipes.component.css']
})
export class MyRecipesComponent {
  recipes: Recipe[] = []

  ngOnInit(): void {
    this.recipesService.getRecipes().subscribe((response: Recipe[]) => {
      this.recipes = response
    })
    
  }

  openRecipe(id: number) {
    this.router.navigateByUrl(`rezept/${id}`)
  }

  constructor(private recipesService: RecipesService, private router: Router) { }

}

