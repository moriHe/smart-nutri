import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Api } from 'api';
import { Observable, map } from 'rxjs';
import { FullRecipe, ShallowRecipe } from './recipes.interface';



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

  constructor(
    private http: HttpClient,
    private api: Api
  ) {}

}
