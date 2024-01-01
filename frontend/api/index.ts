import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { FullRecipe, RecipeIngredientBody, ShallowRecipe, RecipeBody } from "./recipes/recipes.interface";
import { Auth, idToken, authState, User } from '@angular/fire/auth';
import { of, switchMap, take } from "rxjs";

type DbUser = {
    id: number,
    activeFamilyId: number | null
}

export interface Response<T> {
    data: T
    // todo return status from go. Not working right now
    status: number
}

@Injectable({
    providedIn: 'root'
  })
export class Api {
    authState$ = authState(this.auth);
    private user!: DbUser | null
    

    fetchRecipes() {
        return this.http.get<Response<ShallowRecipe[]>>(`http://localhost:8080/familys/${this.user?.activeFamilyId}/recipes`)
    }

    fetchRecipe(id: number) {
        return this.http.get<Response<FullRecipe>>(`http://localhost:8080/recipes/${id}`)
    }

    postRecipe(body: RecipeBody) {
        return this.http.post<Response<{id: number}>>(`http://localhost:8080/familys/${this.user?.activeFamilyId}/recipes`, body)
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

    fetchUser() {
        return this.http.get<Response<DbUser>>("http://localhost:8080/user")
    }

    postUser(fireUid: string) {
        return this.http.post<Response<{userId: number}>>("http://localhost:8080/user", {
            fireUid
        })
    }

    constructor(
        private auth: Auth,
        private http: HttpClient
      ) {
        this.authState$.pipe(
            take(1),
            switchMap((authUser) => {
            if (authUser) {
              return this.fetchUser()
            }
            return of(null)
            })).subscribe((response) => {
                if (response) {
                    this.user = response?.data
                }
            })
      }
}