import { DatePipe } from '@angular/common';
import { ChangeDetectorRef, Component, Inject, LOCALE_ID } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Mealplan, PostMealplanPayload } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Meals, RecipeWithoutIngredients } from 'api/recipes/recipes.interface';
import { Subscription, take } from 'rxjs';
import { MealsService } from 'services/meals.service';
import { CreateMealplanDialogComponent } from '../create-mealplan-dialog/create-mealplan-dialog.component';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { RecipesService } from 'api/recipes/recipes.service';
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import { CreateMealplanBottomsheetComponent } from '../create-mealplan-bottomsheet/create-mealplan-bottomsheet.component';

@Component({
  selector: 'app-mealplans',
  templateUrl: './mealplans.component.html',
  styleUrls: ['./mealplans.component.css'],
})
export class MealplansComponent {
  isMobile = false
  isMobileDialogOpen = false
  mobileSearchQuery = ""
  private recipesSubscription!: Subscription
  recipes: RecipeWithoutIngredients[] = []
  
  selectedMeal!: Meals
  mealplan: Mealplan = []

  private mealplanSubscription!: Subscription
  today: Date = new Date()
  selectedDate: Date = this.today
  private updateMealplan(): void {
      if (this.mealplanSubscription) {
        this.mealplanSubscription.unsubscribe();
      }
  
      this.mealplanSubscription = this.mealplanService
        .getMealplan(this.selectedDate.toISOString())
        .subscribe((response: Mealplan) => {
          if (response) {
            this.mealplan = response;
          }
        });
    }

  previousDay(): void {
    this.selectedDate.setDate(this.selectedDate.getDate() - 1);
    this.updateMealplan()
    this.cdr.detectChanges()
  }

  nextDay(): void {
    this.selectedDate.setDate(this.selectedDate.getDate() + 1);
    this.updateMealplan()
    this.cdr.detectChanges()
  }

  displayDate(): string {

    const yesterday = new Date(this.today);
    yesterday.setDate(this.today.getDate() - 1);
    const tomorrow = new Date(this.today);
    tomorrow.setDate(this.today.getDate() + 1);
    let format = 'EEEE, dd. MMMM';

    if (this.selectedDate.toDateString() === this.today.toDateString()) {
      format = "'Heute', " + format;
    } else if (this.selectedDate.toDateString() === yesterday.toDateString()) {
      format = "'Gestern', " + format;
    } else if (this.selectedDate.toDateString() === tomorrow.toDateString()) {
      format = "'Morgen', " + format;
    }

    const displayDate = this.datePipe.transform(this.selectedDate, format, undefined, this.locale);
    return displayDate || '';
  }
  addMealPlanItem = ({recipeId, meal, portions}: Omit<PostMealplanPayload, "date" | "isShoppingListItem">) =>{
    this.mealplanService.addMealplanItem({
      recipeId,
      meal,
      portions,
      date: this.selectedDate.toISOString(),
      isShoppingListItem: false

    }).pipe(take(1)).subscribe({
      next: () => {
        this.updateMealplan()
        this.isMobileDialogOpen = false
      }
    })
  }

  openDialog = (mealKey: Meals) => {
    if (this.isMobile) {
      this.selectedMeal = mealKey
      this.isMobileDialogOpen = true
      return
    }
    const dialogRef = this.dialog.open(CreateMealplanDialogComponent, 
      {height: "80vh", width: "60vw", data: {selectedMeal: mealKey}});

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.addMealPlanItem({...result, meal: mealKey})
      }
  })
}

  getMealplanForMealType(mealKey: Meals): Mealplan {
    return this.mealplan.filter(item => item.meal === Meals[mealKey]);
  }

  closeMobileDialog() {
    this.isMobileDialogOpen = false
    this.mobileSearchQuery = ""
  }

  searchQueryRecipes() {
    if (this.mobileSearchQuery == "") {
      return this.recipes
    }

    return this.recipes.filter((recipe) => {
      return recipe.name.toLowerCase().includes(this.mobileSearchQuery.toLowerCase())
    })
  }

  openBottomSheet(id: number) {
    const selectedRecipe = this.recipes.find((recipe) => recipe.id === id)
    if (!selectedRecipe) {
      return
    }
    this._bottomSheet.open(CreateMealplanBottomsheetComponent, 
      { data: {...selectedRecipe, addMealplanItem: this.addMealPlanItem, selectedMeal: this.selectedMeal} }
      );
  }

  ngOnInit(): void {
    this.recipesSubscription = this.recipesService.getRecipes().subscribe((response: RecipeWithoutIngredients[]) => {
      this.recipes = response
    })
  

  this.updateMealplan()

  this.breakpointObserver.observe([
    Breakpoints.Handset,
    Breakpoints.Tablet,
  ]).subscribe(result => {
    this.isMobile = result.matches;
  });
}

  ngOnDestroy(): void {
    if (this.mealplanSubscription) {
      this.mealplanSubscription.unsubscribe()
    }
    if (this.recipesSubscription) {
      this.recipesSubscription.unsubscribe();
    }
  }

  constructor(
    private mealplanService: MealplansService,
    private cdr: ChangeDetectorRef,
    private datePipe: DatePipe, 
    @Inject(LOCALE_ID) private locale: string,
    public mealsService: MealsService,
    public dialog: MatDialog,
    private router: Router,
    private breakpointObserver: BreakpointObserver,
    private recipesService: RecipesService,
    private _bottomSheet: MatBottomSheet, 
    ) {}
}
