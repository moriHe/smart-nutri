import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FullRecipe } from 'api/recipes/recipes.interface';
import { RecipesService } from 'api/recipes/recipes.service';
import { MealsService } from 'services/meals.service';

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
        console.log(response)
        this.recipe = response
      });
    })
  }

  constructor(
    private route: ActivatedRoute,
    private recipesService: RecipesService, 
    public mealsService: MealsService
    ) { }

}
