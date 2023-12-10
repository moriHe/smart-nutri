import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';

export type Recipe = {
  id: number,
  name: string
}

@Injectable({
  providedIn: 'root'
})
export class RecipesService {
  getRecipes(): Observable<Recipe[]> {
    return this.http.get<{data: Recipe[]}>("http://localhost:8080/familys/1/recipes").pipe(
      map((response: {data: Recipe[]}) => response.data)
    );
  }

  constructor(
    private http: HttpClient
  ) {}

}
