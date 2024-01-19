import { Component, Input } from '@angular/core';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Markets } from 'api/recipes/recipes.interface';
import { ShoppingListItems } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { isSameDay } from 'date-fns';
import { take } from 'rxjs';
import { MarketsService } from 'services/markets.service';
import { UnitsService } from 'services/units.service';

type ItemsMarkets = {
  TODAY: ShoppingListItems
  [Markets.NONE]: ShoppingListItems
  [Markets.REWE]: ShoppingListItems
  [Markets.EDEKA]: ShoppingListItems
  [Markets.BIO_COMPANY]: ShoppingListItems
  [Markets.WEEKLY_MARKET]: ShoppingListItems
  [Markets.ALDI]: ShoppingListItems
  [Markets.LIDL]: ShoppingListItems
}

type ItemMarketKeys = keyof ItemsMarkets


@Component({
  selector: 'app-shopping-list-view',
  templateUrl: './shopping-list-view.component.html',
  styleUrls: ['./shopping-list-view.component.css']
})
export class ShoppingListViewComponent {
  @Input() openAddItemsView!: () => void
  shoppingList: ShoppingListItems = []
  mealplanNumberNotOnShoppingList = 0


  sortShoppingListItems(): {key: ItemMarketKeys, list: ShoppingListItems}[] {
    const today = new Date()
    const itemsToday = this.shoppingList.filter((item) => {
      return isSameDay(new Date(item.mealplanItem.date), today)
    })
  
    const itemsMarkets: ItemsMarkets = this.shoppingList.reduce((acc, curr) => {
      if (isSameDay(new Date(curr.mealplanItem.date), today)) {
        return acc
      }
      return {
        ...acc,
        [curr.market]: [...acc[curr.market], curr]
      }
    }, {TODAY: itemsToday, NONE: [], REWE: [], EDEKA: [], BIO_COMPANY: [], WEEKLY_MARKET: [], ALDI: [], LIDL: []})
  
    return [{
      key: "TODAY",
      list: itemsToday
    },
    {
      key: Markets.NONE,
      list: itemsMarkets.NONE,
    },
    {
      key: Markets.REWE,
      list: itemsMarkets.REWE,
    },
    {
      key: Markets.EDEKA,
      list: itemsMarkets.EDEKA,
    },
    {
      key: Markets.BIO_COMPANY,
      list: itemsMarkets.BIO_COMPANY,
    },
    {
      key: Markets.WEEKLY_MARKET,
      list: itemsMarkets.WEEKLY_MARKET,
    },
    {
      key: Markets.ALDI,
      list: itemsMarkets.ALDI,
    },
    {
      key: Markets.LIDL,
      list: itemsMarkets.LIDL,
    }
  
  ]
  }

  getAmount(amountPerPortion: number, portions: number) {
    return Math.round(amountPerPortion * portions)
  }


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
          this.shoppingList = response
        }
      })
  }
  ngOnInit(): void {
    this.updateListItems()
  }

  constructor(
    private mealplansService: MealplansService,
    private shoppingListService: ShoppingListService,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}

}
