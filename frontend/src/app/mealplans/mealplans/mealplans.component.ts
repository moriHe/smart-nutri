import { DatePipe } from '@angular/common';
import { ChangeDetectorRef, Component, Inject, LOCALE_ID } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Mealplan } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Meals } from 'api/recipes/recipes.interface';
import { Subscription, take } from 'rxjs';
import { MealsService } from 'services/meals.service';
import { CreateMealplanDialogComponent } from '../create-mealplan-dialog/create-mealplan-dialog.component';

@Component({
  selector: 'app-mealplans',
  templateUrl: './mealplans.component.html',
  styleUrls: ['./mealplans.component.css'],
})
export class MealplansComponent {
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

    ngOnInit(): void {
      this.updateMealplan()
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

  addMealPlanItem() {
    this.mealplanService.addMealplanItem().pipe(take(1)).subscribe({
      next: () => {
        this.updateMealplan()
      }
    })
  }

  openDialog(mealKey: Meals) {
    const dialogRef = this.dialog.open(CreateMealplanDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      // TODO add error if something is missing
      console.log("test")
  })
}
  getMealplanForMealType(mealKey: Meals): Mealplan {
    console.log(this.mealplan)
    console.log(this.mealplan.filter(item => item.meal === Meals[mealKey]))
    return this.mealplan.filter(item => item.meal === Meals[mealKey]);
  }

  ngOnDestroy(): void {
    if (this.mealplanSubscription) {
      this.mealplanSubscription.unsubscribe()
    }
  }

  constructor(
    private mealplanService: MealplansService,
    private cdr: ChangeDetectorRef,
    private datePipe: DatePipe, 
    @Inject(LOCALE_ID) private locale: string,
    public mealsService: MealsService,
    public dialog: MatDialog,
    private router: Router
    ) {}
}
