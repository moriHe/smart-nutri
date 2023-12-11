import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Recipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';

@Component({
  selector: 'app-recipe-details',
  templateUrl: './recipe-details.component.html',
  styleUrls: ['./recipe-details.component.css']
})
export class RecipeDetailsComponent {
  recipe?: Recipe

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      const id = params['id']
      this.recipesService.getRecipe(id).subscribe((response: Recipe) => {
        console.log(response)
        this.recipe = response
      });
    })
  }

  constructor(private recipesService: RecipesService, private route: ActivatedRoute) { }

}
