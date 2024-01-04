import { Injectable } from "@angular/core";
import { FullRecipe, RecipeBody, RecipeIngredientBody, ShallowRecipe } from "./recipes.interface";
import { Response } from "api";
import { HttpClient } from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class RecipesEndpointsService {

  fetchRecipes() {
    return this.http.get<Response<ShallowRecipe[]>>(`http://localhost:8080/recipes`);
}

fetchRecipe(id: number) {
    return this.http.get<Response<FullRecipe>>(`http://localhost:8080/recipes/${id}`)
}

postRecipe(body: RecipeBody) {
    return this.http.post<Response<{ id: number }>>(
      `http://localhost:8080/recipes`,
      body
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
    private http: HttpClient
  ) {}
}