import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { FullRecipe, RecipeIngredientBody, RecipeIngredient, ShallowRecipe } from "./recipes/recipes.interface";
import { catchError } from "rxjs";

interface Response<T> {
    data: T
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

    postRecipeIngredient(recipeId: number, body: RecipeIngredientBody) {
        return this.http.post<Response<string>>(`http://localhost:8080/recipes/${recipeId}/recipeingredient`, body)
    }

    constructor(
        private http: HttpClient
      ) {}
}