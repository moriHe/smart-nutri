import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Api } from 'api';
import { Observable, map } from 'rxjs';
import { Recipe } from './recipes.interface';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {
  getRecipes(): Observable<Recipe[]> {
    return this.api.fetchRecipes().pipe(
      map((response: {data: Recipe[]}) => response.data)
    );
  }

  getRecipe(id: number): Observable<Recipe>{
    return this.api.fetchRecipe(id).pipe(map((response: {data: Recipe}) => response.data))
  }

  constructor(
    private http: HttpClient,
    private api: Api
  ) {}

}
