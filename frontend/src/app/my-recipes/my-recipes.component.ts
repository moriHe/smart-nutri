import { Component } from '@angular/core';
import { Recipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';

@Component({
  selector: 'app-my-recipes',
  templateUrl: './my-recipes.component.html',
  styleUrls: ['./my-recipes.component.css']
})
export class MyRecipesComponent {
  constructor(private recipesService: RecipesService) { }

  recipes: Recipe[] = []

  ngOnInit(): void {
    this.recipesService.getRecipes().subscribe((response: Recipe[]) => {
      this.recipes = response
    })
    
  }
}

