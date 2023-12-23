import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { FullRecipe, RecipeIngredientBody, ShallowRecipe, RecipeBody } from "./recipes/recipes.interface";

export interface Response<T> {
    data: T
    // todo return status from go. Not working right now
    status: number
}

@Injectable({
    providedIn: 'root'
  })
export class Api {
    fetchRecipes() {
        return this.http.get<Response<ShallowRecipe[]>>("http://localhost:8080/familys/1/recipes")
    }

    fetchRecipe(id: number) {
        return this.http.get<Response<FullRecipe>>(`http://localhost:8080/recipes/${id}`)
    }

    postRecipe(body: RecipeBody) {
        return this.http.post<Response<{id: number}>>(`http://localhost:8080/familys/1/recipes`, body)
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