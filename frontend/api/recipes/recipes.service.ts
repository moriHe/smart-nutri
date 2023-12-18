import { Injectable } from '@angular/core';
import { Api } from 'api';
import { Observable, map } from 'rxjs';
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from './recipes.interface';



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

  addRecipe(body: RecipeBody): Observable<number>{
    return this.api.postRecipe(body).pipe(map((response: {status: number}) => response.status))
  }

  addRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<string>{
    return this.api.postRecipeIngredient(recipeId, body).pipe(map((response: {data: string}) => response.data))
  }

  constructor(private api: Api) {}

}
