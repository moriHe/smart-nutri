import { Injectable } from "@angular/core";
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from "./recipes.interface";
import { Response } from "api";
import { HttpClient } from "@angular/common/http";
import { DbUser, UserService } from "api/user/user.service";
import { Observable, map, of, switchMap } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class RecipesEndpointsService {

  fetchRecipes() {
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
        })
      );
}

fetchRecipe(id: number) {
    return this.http.get<Response<FullRecipe>>(`http://localhost:8080/recipes/${id}`)
}

postRecipe(body: RecipeBody) {
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

deleteRecipe(id: number) {
    return this.http.delete<Response<string>>(`http://localhost:8080/recipes/${id}`)
}

postRecipeIngredient(recipeId: number, body: RecipeIngredientBody) {
    return this.http.post<Response<number>>(`http://localhost:8080/recipes/${recipeId}/recipeingredient`, body)
}

deleteRecipeIngredient(ingredientId: number) {
    return this.http.delete<Response<string>>(`http://localhost:8080/recipes/recipeingredient/${ingredientId}`)
}

constructor(
    private userService: UserService,
    private http: HttpClient
  ) {}
}


// TODO Streamline api for recipes and other so it can get more generic
// user/:userId/familys/:familyId
// import { Observable, of } from 'rxjs';
// import { switchMap } from 'rxjs/operators';
// import { HttpClient } from '@angular/common/http';
// import { Response } from 'api';
// import { DbUser, UserService } from 'api/user/user.service';
// import { ShallowRecipe } from './recipes.interface';

// @Injectable({
//   providedIn: 'root'
// })
// export class GenericService {
//   constructor(private userService: UserService, private http: HttpClient) {}

//   fetchData<T>(url: string): Observable<Response<T> | null> {
//     return this.userService.user.pipe(
//       switchMap((user) => {
//         if (user) {
//           const apiUrl = `http://localhost:8080/familys/${user.data.activeFamilyId}${url}`;
//           return this.http.get<Response<T>>(apiUrl);
//         } else {
//           return of(null);
//         }
//       })
//     );
//   }

//   // Example usage for fetching recipes
//   fetchRecipes(): Observable<Response<ShallowRecipe[]>> {
//     return this.fetchData<ShallowRecipe[]>('/recipes');
//   }
// }