import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Observable, Subscription, map, of, switchMap } from 'rxjs';
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from './recipes.interface';
import { HttpClient } from '@angular/common/http';
import { DbUser, UserService } from 'api/user/user.service';
import { RecipesEndpointsService } from './recipes.endpoints.service';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {


  getRecipes(): Observable<ShallowRecipe[]> {
    return this.recipesEndpointService.fetchRecipes()
  }

  getRecipe(id: number): Observable<FullRecipe>{
    return this.recipesEndpointService.fetchRecipe(id)
  }

  addRecipe(body: RecipeBody): Observable<Response<{id: number}> | null>{
    return this.recipesEndpointService.postRecipe(body)
  }

  removeRecipe(id: number): Observable<string>{
    return this.recipesEndpointService.deleteRecipe(id)
  }

  addRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<number>{
    return this.recipesEndpointService.postRecipeIngredient(recipeId, body)
  }

  removeRecipeIngredient(recipeIngredientId: number): Observable<string>{
    return this.recipesEndpointService.deleteRecipeIngredient(recipeIngredientId)
  }

  constructor(
    private recipesEndpointService: RecipesEndpointsService,
    ) {}

}
