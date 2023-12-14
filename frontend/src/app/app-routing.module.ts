import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HeroesComponent } from './heroes/heroes.component';
import { MyRecipesComponent } from './my-recipes/my-recipes.component';
import { RecipeDetailsComponent } from './recipe-details/recipe-details.component';
import { SearchComponent } from './search/search.component';

const routes: Routes = [
  { path: 'heroes', component: HeroesComponent },
  {path: "meine-rezepte", component: MyRecipesComponent},
  {path: "rezept/:id", component: RecipeDetailsComponent},
  {path: "suche", component: SearchComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
