import { Component } from '@angular/core';
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
        console.log(response)
        this.recipe = response
      });
    })
  }

  openSearch() {
    if (this.recipe) {
      this.router.navigateByUrl(`suche/${this.recipe.id}`)
    }
  }

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private recipesService: RecipesService, 
    public mealsService: MealsService,
    public unitsService: UnitsService,
    public marketsService: MarketsService
    ) { }

}
