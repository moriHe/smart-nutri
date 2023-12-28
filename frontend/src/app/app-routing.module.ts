import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HeroesComponent } from './heroes/heroes.component';
import { MyRecipesComponent } from './my-recipes/my-recipes.component';
import { RecipeDetailsComponent } from './recipe-details/recipe-details.component';
import { SearchComponent } from './search/search.component';
import { SignupComponent } from './signup/signup.component';
import { SignupSuccessComponent } from './signup-success/signup-success.component';

const routes: Routes = [
  { path: 'heroes', component: HeroesComponent },
  {path: "meine-rezepte", component: MyRecipesComponent},
  {path: "rezept/:id", component: RecipeDetailsComponent},
  {path: "suche/:recipeId", component: SearchComponent},
  {path: "registrieren", component: SignupComponent},
  {path: "willkommen", component: SignupSuccessComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
