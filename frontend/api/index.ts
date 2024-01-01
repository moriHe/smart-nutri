// import { HttpClient } from "@angular/common/http";
// import { Injectable } from "@angular/core";
// import { FullRecipe, RecipeIngredientBody, ShallowRecipe, RecipeBody } from "./recipes/recipes.interface";
// import { Auth, idToken, authState, User } from '@angular/fire/auth';
// import { of, switchMap, take } from "rxjs";

// type DbUser = {
//     id: number,
//     activeFamilyId: number | null
// }

export interface Response<T> {
    data: T
    // todo return status from go. Not working right now
    status: number
}

// @Injectable({
//     providedIn: 'root'
//   })
// export class Api {
//     authState$ = authState(this.auth);
//     private user!: DbUser | null
//     userSubscription = this.authState$.pipe(
//         take(1),
//         switchMap((authUser) => {
//         if (authUser) {
//           return this.fetchUser()
//         }
//         return of(null)
//         })).subscribe((response) => {
//             console.log("test")
//             if (response) {
//                 this.user = response?.data
//             }
//         })
    
//         ngOnDestroy(): void {
//             this.userSubscription.unsubscribe()
//         }

//     fetchRecipes() {
//         return 
//     }

//     fetchRecipe(id: number) {
//         return 
//     }

//     postRecipe(body: RecipeBody) {
//         return 
//     }

//     deleteRecipe(id: number) {
//         return 
//     }

//     postRecipeIngredient(recipeId: number, body: RecipeIngredientBody) {
//         return 
//     }

//     deleteRecipeIngredient(ingredientId: number) {
//         return 
//     }

//     fetchUser() {
//         return 
//     }

//     postUser(fireUid: string) {
//         return this.http.post<Response<{userId: number}>>("http://localhost:8080/user", {
//             fireUid
//         })
//     }

//     constructor(
//         private auth: Auth,
//         private http: HttpClient
//       ) {}
// }