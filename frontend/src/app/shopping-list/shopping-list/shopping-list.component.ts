import { Component } from '@angular/core';
import { FullMealplanItem, Mealplans } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Markets } from 'api/recipes/recipes.interface';
import { AddToShoppingList, RecipeIngredientItem, ShoppingListItems } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
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
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  isShoppingListView = true
  mealplansNotOnShoppingList: Mealplans = []
  shoppingList: ShoppingListItems = []
  categories: (Markets | "TODAY")[] = ["TODAY", ...Object.values(Markets)]
  mealplanItemInView?: FullMealplanItem
  addToShoppingList: AddToShoppingList[] = []

  
  // TODO: Add mealplan items to shopping list
  // *** kind of a session approach (abbrechen, speichern und weiter, speichern und beenden)
  // display shopping list items
  // remove shopping list items

  removeItem(recipeIngredientId: number) {
    if (this.mealplanItemInView) {
      this.mealplanItemInView = {
        ...this.mealplanItemInView,
        recipe: {
          ...this.mealplanItemInView.recipe,
          recipeIngredients: this.mealplanItemInView.recipe.recipeIngredients.filter(
            (ingredient: RecipeIngredientItem) => ingredient.id != recipeIngredientId)
        }
      }
    }
  }

openAddItemsView() {
  this.mealplansService.getMealplanItem(this.mealplansNotOnShoppingList[0].id).pipe(take(1))
  .subscribe((response) => {
    console.log(response)
    // todo make this proper
    this.mealplanItemInView = response
  })
  this.isShoppingListView = false
}

openShoppingListView() {
  this.isShoppingListView = true
  this.mealplanItemInView = undefined
}

addMealplanToShoppingList(action: "finish" | "next") {
  let updatedAddToShoppingList: AddToShoppingList[] = this.addToShoppingList
  if (this.mealplanItemInView) {
    const shoppingListItems: AddToShoppingList[] = this.mealplanItemInView.recipe.recipeIngredients.map(
      (item) => {
        return {
          recipeIngredientId: item.id,
          market: item.market,
          isBio: item.isBio
        }
      }
    )
    updatedAddToShoppingList = [...updatedAddToShoppingList, ...shoppingListItems]
  }
  if (action === "finish") {
    this.shoppingListService.addShoppingListItems(updatedAddToShoppingList, this.mealplanItemInView!.id).pipe(take(1)).subscribe((response) => {
      if (response) {
        this.mealplanItemInView = undefined
        this.isShoppingListView = true
        this.updateListItems()
      }
    })
  }
 
}
// MarketDisplay = {
//   [Markets.NONE]: "-",
//   [Markets.REWE]: "Rewe",
//   [Markets.EDEKA]: "Edeka",
//   [Markets.BIO_COMPANY]: "Bio Company",
//   [Markets.WEEKLY_MARKET]: "Wochenmarkt",
//   [Markets.ALDI]: "Aldi",
//   [Markets.LIDL]: "Lidl"
// }

sortShoppingListItems(): {key: ItemMarketKeys, list: ShoppingListItems}[] {
  const today = new Date()
  const itemsToday = this.shoppingList.filter((item) => {
    return new Date(item.mealplanItem.date) === today
  })

  const itemsMarkets: ItemsMarkets = this.shoppingList.reduce((acc, curr) => {
    if (new Date(curr.mealplanItem.date) === today) {
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
  console.log(id)
}

updateListItems() {
  this.mealplansService.getMealplan({date: new Date().toISOString(), forShoppingList: true})
    .pipe(take(1)).subscribe((response) => {
      if (response) {
        this.mealplansNotOnShoppingList = response
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
    private shoppingListService: ShoppingListService,
    private mealplansService: MealplansService,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}
}
