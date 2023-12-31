import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HeroesComponent } from './heroes/heroes.component';
import { MyRecipesComponent } from './my-recipes/my-recipes.component';
import { RecipeDetailsComponent } from './recipe-details/recipe-details.component';
import { SearchComponent } from './search/search.component';
import { SignupComponent } from './signup/signup.component';
import { SignupSuccessComponent } from './signup-success/signup-success.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { fireAuthGuard } from 'api/user/fire-auth.guard';

const routes: Routes = [
  { path: 'heroes', component: HeroesComponent },
  {path: "", component: LandingPageComponent},
  {path: "meine-rezepte", component: MyRecipesComponent, canActivate: [fireAuthGuard]},
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
