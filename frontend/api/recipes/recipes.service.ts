import { Injectable } from '@angular/core';
import { Observable, map, take } from 'rxjs';
import { FullRecipe, RecipeBody, RecipeIngredientBody, RecipeWithoutIngredients } from './recipes.interface';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment.development'
import { Response } from 'api';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {


  getRecipes(): Observable<RecipeWithoutIngredients[]> {
    return  this.http.get<Response<RecipeWithoutIngredients[]>>(`${environment.backendBaseUrl}/recipes`)
    .pipe(
      map((response: { data: RecipeWithoutIngredients[] }) => response ? response.data : [])
    )
  }

  getRecipe(id: number): Observable<FullRecipe>{
    return this.http.get<Response<FullRecipe>>(`${environment.backendBaseUrl}/recipes/${id}`)
    .pipe
    (map((response: Response<FullRecipe>) => response.data)
    )
  }

  addRecipe(body: RecipeBody): Observable<Response<{id: number}> | null>{
    return this.http.post<Response<{ id: number }>>(
      `${environment.backendBaseUrl}/recipes`,
      body
    )
  }

  removeRecipe(id: number): Observable<string>{
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/recipes/${id}`)
    .pipe(
      take(1), 
      map((response: Response<string>) => response.data)
    )
  }

  addRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<number>{
    return this.http.post<Response<number>>(`${environment.backendBaseUrl}/recipes/${recipeId}/recipeingredient`, body)
    .pipe(
      map((response: Response<number>) => response.data)
      )
  }

  removeRecipeIngredient(recipeIngredientId: number): Observable<string>{
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/recipes/recipeingredient/${recipeIngredientId}`)
    .pipe(
      map((response: Response<string>) => response.data)
      )
  }

  constructor(
    private http: HttpClient,
    ) {}

}
