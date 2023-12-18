import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Api } from 'api';
import { Observable, map } from 'rxjs';
import { FullRecipe, RecipeIngredientBody, RecipeIngredient, ShallowRecipe } from './recipes.interface';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {
  getRecipes(): Observable<ShallowRecipe[]> {
    return this.api.fetchRecipes().pipe(
      map((response: {data: ShallowRecipe[]}) => response.data)
    );
  }

  getRecipe(id: number): Observable<FullRecipe>{
    return this.api.fetchRecipe(id).pipe(map((response: {data: FullRecipe}) => response.data))
  }

  postRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<string>{
    return this.api.postRecipeIngredient(recipeId, body).pipe(map((response: {data: string}) => response.data))
  }

  constructor(
    private http: HttpClient,
    private api: Api
  ) {}

}
