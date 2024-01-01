import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Observable, Subscription, map, of, switchMap } from 'rxjs';
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from './recipes.interface';
import { HttpClient } from '@angular/common/http';
import { DbUser, UserService } from 'api/user/user.service';



@Injectable({
  providedIn: 'root'
})
export class RecipesService {


  getRecipes(): Observable<ShallowRecipe[]> {
    return this.userService.user.pipe(
      // Wait for user data to be available
      switchMap((user) => {
        if (user) {
          // Make the HTTP request with the user's activeFamilyId
          return this.http.get<Response<ShallowRecipe[]>>(`http://localhost:8080/familys/${user.data.activeFamilyId}/recipes`);
        } else {
          // Handle the case when user data is not available
          return of(null);
        }
      }),
      map((response: { data: ShallowRecipe[] } | null) => response ? response.data : [])
    );
  }

  getRecipe(id: number): Observable<FullRecipe>{
    return this.http.get<Response<FullRecipe>>(`http://localhost:8080/recipes/${id}`).pipe(map((response: Response<FullRecipe>) => response.data))
  }

  addRecipe(body: RecipeBody): Observable<Response<{id: number}> | null>{
    return this.userService.user.pipe(
      switchMap((user) => {
        // Check if user is available
        if (user) {
          // Make the HTTP request using user data
          return this.http.post<Response<{ id: number }>>(
            `http://localhost:8080/familys/${user.data.activeFamilyId}/recipes`,
            body
          );
        } else {
          // If user is not available, return an empty observable or handle it based on your use case
          return of(null); // You can use of(null) or throwError() depending on your needs
        }
      })
    );
  }

  removeRecipe(id: number): Observable<string>{
    return this.http.delete<Response<string>>(`http://localhost:8080/recipes/${id}`).pipe(map((response: Response<string>) => response.data))
  }

  addRecipeIngredient(recipeId: number, body: RecipeIngredientBody): Observable<number>{
    return this.http.post<Response<number>>(`http://localhost:8080/recipes/${recipeId}/recipeingredient`, body).pipe(map((response: Response<number>) => response.data))
  }

  removeRecipeIngredient(recipeIngredientId: number): Observable<string>{
    return this.http.delete<Response<string>>(`http://localhost:8080/recipes/recipeingredient/${recipeIngredientId}`).pipe(map((response: Response<string>) => response.data))
  }

  constructor(
    private userService: UserService,
    private http: HttpClient
    ) {}

}
