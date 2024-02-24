import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { DatePipe } from '@angular/common';
import { Component, Inject, Input, LOCALE_ID, Renderer2, ViewChild } from '@angular/core';
import { MatExpansionPanel } from '@angular/material/expansion';
import { MatSnackBarRef, SimpleSnackBar } from '@angular/material/snack-bar';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Markets } from 'api/recipes/recipes.interface';
import { ShoppingListByCategory, shoppingListCategories } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { take } from 'rxjs';
import { MarketsService } from 'services/markets.service';
import { SnackbarService } from 'services/snackbar.service';
import { UnitsService } from 'services/units.service';

@Component({
  selector: 'app-shopping-list-view',
  templateUrl: './shopping-list-view.component.html',
  styleUrls: ['./shopping-list-view.component.css']
})
export class ShoppingListViewComponent {
  isMobile = false
  snackbarRef!: MatSnackBarRef<SimpleSnackBar>
  @Input() openAddItemsView!: () => void
  categories = shoppingListCategories
  shoppingListByCategory: ShoppingListByCategory = {
    TODAY: [], REWE: [], EDEKA: [], BIO_COMPANY: [], WEEKLY_MARKET: [], ALDI: [], LIDL: [], NONE: []
  }
  mealplanNumberNotOnShoppingList = 0
  @ViewChild(MatExpansionPanel, {static: true}) matExpansionPanelElement!: MatExpansionPanel;


  closePanel() {
    this.matExpansionPanelElement.close()
}

  openPanel() {
    this.matExpansionPanelElement.open()
  }

  togglePanel() {
    this.matExpansionPanelElement.toggle()
  }

displayDate(dateString: string) {
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
  this.snackbarRef = this.snackbarService.openSnackbar(displayDate || "", "")
}

displayMarket(market: Markets) {
  console.log("test")
  this.snackbarRef = this.snackbarService.openSnackbar(this.marketsService.MarketDisplay[market], "")
}

dismissSnackbar() {
  this.snackbarRef.dismiss()
}

roundAmount(portionAmount: number, amountPerPortion: number): string {
  const roundedAmount = (portionAmount * amountPerPortion).toFixed(1)
  return roundedAmount.endsWith(".0") ? parseInt(roundedAmount).toString() : roundedAmount
}

removeFromShoppingList(ids: number[], event: Event) {
  event.stopPropagation()
  this.shoppingListService.newRemoveShoppingListItems(ids).pipe(take(1)).subscribe((response) => {
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
        this.shoppingListByCategory = response
      }
      })
  }

  ngOnInit(): void {
    this.updateListItems()

    this.breakpointObserver.observe([
      Breakpoints.Handset,
      Breakpoints.Tablet,
    ]).subscribe(result => {
      this.isMobile = result.matches;
    });
  }

  constructor(
    private snackbarService: SnackbarService,
    private breakpointObserver: BreakpointObserver,
    private datePipe: DatePipe,
    @Inject(LOCALE_ID) private locale: string,
    private mealplansService: MealplansService,
    private shoppingListService: ShoppingListService,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}

}
