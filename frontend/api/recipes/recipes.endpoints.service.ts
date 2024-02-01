import { Injectable } from "@angular/core";
import { FullRecipe, RecipeBody, RecipeIngredientBody, RecipeWithoutIngredients } from "./recipes.interface";
import { Response } from "api";
import { HttpClient } from "@angular/common/http";
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class RecipesEndpointsService {

  fetchRecipes() {
    return this.http.get<Response<RecipeWithoutIngredients[]>>(`${environment.backendBaseUrl}/recipes`);
}

fetchRecipe(id: number) {
    return this.http.get<Response<FullRecipe>>(`${environment.backendBaseUrl}/recipes/${id}`)
}

postRecipe(body: RecipeBody) {
    return this.http.post<Response<{ id: number }>>(
      `${environment.backendBaseUrl}/recipes`,
      body
    );
}

deleteRecipe(id: number) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/recipes/${id}`)
}

postRecipeIngredient(recipeId: number, body: RecipeIngredientBody) {
    return this.http.post<Response<number>>(`${environment.backendBaseUrl}/recipes/${recipeId}/recipeingredient`, body)
}

deleteRecipeIngredient(ingredientId: number) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/recipes/recipeingredient/${ingredientId}`)
}

constructor(
    private http: HttpClient
  ) {}
}
