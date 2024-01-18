import { NgModule, LOCALE_ID, APP_INITIALIZER } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { HeroesComponent } from './heroes/heroes.component';
import { MyRecipesComponent } from './my-recipes/my-recipes.component';
import { RecipeDetailsComponent } from './recipe-details/recipe-details.component';
import { SearchComponent } from './search/search.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import { SearchIngredientDialogComponent } from './search-ingredient-dialog/search-ingredient-dialog.component';
import { MatDialogModule } from '@angular/material/dialog';
import {MatSelectModule} from '@angular/material/select';
import {MatChipsModule} from '@angular/material/chips';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { CreateRecipeDialogComponent } from './create-recipe-dialog/create-recipe-dialog.component';
import {CdkAccordionModule} from '@angular/cdk/accordion';
import { SignupComponent } from './signup/signup.component';
import { provideFirebaseApp, initializeApp } from '@angular/fire/app';
import { getAuth, provideAuth } from '@angular/fire/auth';
import { environment } from 'src/environments/environment.development';
import { SignupSuccessComponent } from './signup-success/signup-success.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { AuthInterceptor } from './auth.interceptor';
import { MealplansComponent } from './mealplans/mealplans/mealplans.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { CommonModule, DatePipe } from '@angular/common';
import { CreateMealplanDialogComponent } from './mealplans/create-mealplan-dialog/create-mealplan-dialog.component';
import { CreateMealplanBottomsheetComponent } from './mealplans/create-mealplan-bottomsheet/create-mealplan-bottomsheet.component';
import { MatBottomSheetModule } from '@angular/material/bottom-sheet';
import { MealplanCardsComponent } from './mealplans/mealplan-cards/mealplan-cards.component';
import {MatSidenavModule} from '@angular/material/sidenav';
import { NavigationButtonsComponent } from './header/navigation-buttons/navigation-buttons.component';
import { ShoppingListComponent } from './shopping-list/shopping-list/shopping-list.component';
import {MatCardModule} from '@angular/material/card';
import { ShoppingListViewComponent } from './shopping-list/shopping-list-view/shopping-list-view.component';
import { NotOnShoppingListViewComponent } from './shopping-list/not-on-shopping-list-view/not-on-shopping-list-view.component';
import { ShoppingListBellComponent } from './shopping-list/shopping-list-bell/shopping-list-bell.component';
import { SupabaseService } from 'api/supabase.service';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    HeroesComponent,
    MyRecipesComponent,
    RecipeDetailsComponent,
    SearchComponent,
    SearchIngredientDialogComponent,
    CreateRecipeDialogComponent,
    SignupComponent,
    SignupSuccessComponent,
    LandingPageComponent,
    MealplansComponent,
    CreateMealplanDialogComponent,
    CreateMealplanBottomsheetComponent,
    MealplanCardsComponent,
    NavigationButtonsComponent,
    ShoppingListComponent,
    ShoppingListViewComponent,
    NotOnShoppingListViewComponent,
    ShoppingListBellComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatIconModule,
    MatButtonModule,
    ReactiveFormsModule,
    MatDialogModule,
    MatSelectModule,
    MatChipsModule,
    MatSnackBarModule,
    CdkAccordionModule,
    MatDatepickerModule,
    MatNativeDateModule,
    CommonModule,
    MatBottomSheetModule,
    MatSidenavModule,
    MatCardModule,
    provideFirebaseApp(() => initializeApp(environment.firebase)),
    provideAuth(() => getAuth()),
  ],
  providers: [
    SupabaseService,
    {
      provide: APP_INITIALIZER,
      useFactory: (supabaseService: SupabaseService) => () => supabaseService.initialize(),
      deps: [SupabaseService],
      multi: true,
    },
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    },
    {
      provide: LOCALE_ID,
      useValue: 'de'
    },
    DatePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
