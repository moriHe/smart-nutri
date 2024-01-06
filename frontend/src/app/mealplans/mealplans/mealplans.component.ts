import { DatePipe } from '@angular/common';
import { ChangeDetectorRef, Component, Inject, LOCALE_ID } from '@angular/core';
import { Mealplan } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Subscription, take } from 'rxjs';

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
    // send console.log(this.today.toUTCString()) to backend
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

  ngOnDestroy(): void {
    if (this.mealplanSubscription) {
      this.mealplanSubscription.unsubscribe()
    }
  }

  constructor(
    private mealplanService: MealplansService,
    private cdr: ChangeDetectorRef,
    private datePipe: DatePipe, 
    @Inject(LOCALE_ID) private locale: string
    ) {}
}
