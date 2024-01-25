import { DatePipe } from '@angular/common';
import { Component, Inject, Input, LOCALE_ID } from '@angular/core';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Markets } from 'api/recipes/recipes.interface';
import { ShoppingListByCategory, shoppingListCategories } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { isSameDay } from 'date-fns';
import { take } from 'rxjs';
import { MarketsService } from 'services/markets.service';
import { UnitsService } from 'services/units.service';

@Component({
  selector: 'app-shopping-list-view',
  templateUrl: './shopping-list-view.component.html',
  styleUrls: ['./shopping-list-view.component.css']
})
export class ShoppingListViewComponent {
  @Input() openAddItemsView!: () => void
  categories = shoppingListCategories
  shoppingListByCategory: ShoppingListByCategory = {
    TODAY: [], REWE: [], EDEKA: [], BIO_COMPANY: [], WEEKLY_MARKET: [], ALDI: [], LIDL: [], NONE: []
  }
  mealplanNumberNotOnShoppingList = 0

// todo add remove all logic
displayDate(dateString: string): string {
  const mealplanDateStr = new Date(dateString).toDateString()
  const today = new Date()
  const yesterday = new Date(today);
  yesterday.setDate(today.getDate() - 1);
  const tomorrow = new Date(today);
  tomorrow.setDate(today.getDate() + 1);
  let format = 'EE, dd. MMM';

  if (mealplanDateStr === today.toDateString()) {
    format = "'Heute', " + format;
  } else if (mealplanDateStr === yesterday.toDateString()) {
    format = "'Gestern', " + format;
  } else if (mealplanDateStr === tomorrow.toDateString()) {
    format = "'Morgen', " + format;
  }
  const displayDate = this.datePipe.transform(mealplanDateStr, format, undefined, this.locale);
  return displayDate || '';
}

roundAmount(portionAmount: number, amountPerPortion: number): string {
  const roundedAmount = (portionAmount * amountPerPortion).toFixed(1)
  return roundedAmount.endsWith(".0") ? parseInt(roundedAmount).toString() : roundedAmount
}

newRemoveFromShoppingList(ids: number[], event: Event) {
  event.stopPropagation()
  this.shoppingListService.newRemoveShoppingListItems(ids).pipe(take(1)).subscribe((response) => {
    if (response) {
      this.updateListItems()
    }
  })
}
// todo display sub stuff
// todo delete single sub
  removeFromShoppingList(id: number) {
    this.shoppingListService.removeShoppingListItem(id).pipe(take(1)).subscribe((response) => {
      if (response) {
        this.updateListItems()
      }
    })
  }


  updateListItems() {
    this.mealplansService.getMealplan({date: new Date().toISOString(), forShoppingList: true})
      .pipe(take(1)).subscribe((response) => {
        if (response) {
          this.mealplanNumberNotOnShoppingList = response.length
        }
      })

      this.shoppingListService.getShoppingList().pipe(take(1)).subscribe((response) => {
        if (response) {
          console.log(response)
        this.shoppingListByCategory = response
      }
      })
  }
  ngOnInit(): void {
    this.updateListItems()
  }

  constructor(
    private datePipe: DatePipe,
    @Inject(LOCALE_ID) private locale: string,
    private mealplansService: MealplansService,
    private shoppingListService: ShoppingListService,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}

}
