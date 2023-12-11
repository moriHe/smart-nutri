import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { FullRecipe, ShallowRecipe } from "./recipes/recipes.interface";

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

    constructor(
        private http: HttpClient
      ) {}
}