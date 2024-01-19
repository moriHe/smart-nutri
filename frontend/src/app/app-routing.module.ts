import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HeroesComponent } from './heroes/heroes.component';
import { MyRecipesComponent } from './my-recipes/my-recipes.component';
import { RecipeDetailsComponent } from './recipe-details/recipe-details.component';
import { SearchComponent } from './search/search.component';
import { SignupComponent } from './signup/signup.component';
import { SignupSuccessComponent } from './signup-success/signup-success.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { authGuard } from 'guards/auth.guard';
import { noUserGuard } from 'guards/no-user.guard';
import { noFamilyGuard } from 'guards/no-family.guard';
import { familyGuard } from 'guards/family.guard';
import { MealplansComponent } from './mealplans/mealplans/mealplans.component';
import { ShoppingListComponent } from './shopping-list/shopping-list/shopping-list.component';

const routes: Routes = [
  { path: 'heroes', component: HeroesComponent },
  {path: "", component: LandingPageComponent, canActivate: [noUserGuard]},
  {path: "meine-rezepte", component: MyRecipesComponent, canActivate: [authGuard, familyGuard]},
  {path: "rezept/:id", component: RecipeDetailsComponent, canActivate: [authGuard, familyGuard]},
  {path: "suche/:recipeId", component: SearchComponent, canActivate: [authGuard, familyGuard]},
  {path: "registrieren", component: SignupComponent, canActivate: [noUserGuard]},
  {path: "willkommen", component: SignupSuccessComponent, canActivate: [authGuard, noFamilyGuard]},
  {path: "essensplan", component: MealplansComponent, canActivate: [authGuard, familyGuard]},
  {path: "einkaufszettel", component: ShoppingListComponent, canActivate: [authGuard, familyGuard]}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
