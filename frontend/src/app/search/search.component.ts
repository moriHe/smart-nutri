import { Component } from '@angular/core';
import { Result, TypesenseService } from '../typesense.service';
import { SearchResponseHit } from 'typesense/lib/Typesense/Documents';
import { RecipesService } from 'api/recipes/recipes.service';
import { ActivatedRoute } from '@angular/router';
import { MarketsService } from 'services/markets.service';
import { UnitsService } from 'services/units.service';
import { Markets, Units } from 'api/recipes/recipes.interface';





@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  recipeId?: number
  results: SearchResponseHit<Result>[] = []

  ngOnInit(): void {
    this.route.params.subscribe(params => {
     this.recipeId = params['recipeId']
    })
    
   this.typesenseService.search("*").then((res) => {
    if (res) {
      this.results = res
    }
   })
  }

  addIngredient(ingredientId: number) {
    if (this.recipeId) {
      console.log(this.recipeId)
      this.recipesService.postRecipeIngredient(this.recipeId, {
        ingredientId,
        amountPerPortion: 3,
        isBio: false,
        market: Markets.REWE,
        unit: Units.GRAM
      }).subscribe()
    }
  }

  constructor(
    private route: ActivatedRoute,
    private typesenseService: TypesenseService,
    private recipesService: RecipesService
    ) { }


}
