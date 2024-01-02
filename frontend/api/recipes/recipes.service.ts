import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Observable, Subscription, map, of, switchMap } from 'rxjs';
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from './recipes.interface';
import { RecipesEndpointsService } from './recipes.endpoints.service';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {


  getRecipes(): Observable<ShallowRecipe[]> {
    return this.recipesEndpointService.fetchRecipes().pipe(
      map((response: { data: ShallowRecipe[] } | null) => response ? response.data : [])
    )
  }

  getRecipe(id: number): Observable<FullRecipe>{
    return this.recipesEndpointService.fetchRecipe(id).pipe(map((response: Response<FullRecipe>) => response.data))
  }

  addRecipe(body: RecipeBody): Observable<Response<{id: number}> | null>{
    return this.recipesEndpointService.postRecipe(body)
  }

  removeRecipe(id: number): Observable<string>{
    return this.recipesEndpointService.deleteRecipe(id).pipe(map((response: Response<string>) => response.data))
  }

  addRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<number>{
    return this.recipesEndpointService.postRecipeIngredient(recipeId, body).pipe(map((response: Response<number>) => response.data))
  }

  removeRecipeIngredient(recipeIngredientId: number): Observable<string>{
    return this.recipesEndpointService.deleteRecipeIngredient(recipeIngredientId).pipe(map((response: Response<string>) => response.data))
  }

  constructor(
    private recipesEndpointService: RecipesEndpointsService,
    ) {}

}
