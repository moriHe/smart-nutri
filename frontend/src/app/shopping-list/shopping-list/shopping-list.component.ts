import { Component } from '@angular/core';
import { Mealplans } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { Markets } from 'api/recipes/recipes.interface';
import { ShoppingListItems } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { take } from 'rxjs';
import { MarketsService } from 'services/markets.service';

@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  isShoppingListView = true
  mealplansNotOnShoppingList: Mealplans = []
  mealplanItemInView: any = undefined
  shoppingList: ShoppingListItems = []
  markets: Markets[] = Object.values(Markets)
  selectedMarket: Markets = Markets.NONE

  
  // TODO: Add mealplan items to shopping list
  // *** kind of a session approach (abbrechen, speichern und weiter, speichern und beenden)
  // display shopping list items
  // remove shopping list items


openAddItemsView() {
  this.mealplansService.getMealplanItem(this.mealplansNotOnShoppingList[0].id).pipe(take(1))
  .subscribe((response) => {
    console.log(response)
    // todo make this proper
    this.selectedMarket = response.recipe.recipeIngredients[0].market
    this.mealplanItemInView = response
  })
  this.isShoppingListView = false
}

ngOnInit(): void {
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
  constructor(
    private shoppingListService: ShoppingListService,
    private mealplansService: MealplansService,
    public marketsService: MarketsService
  ) {}
}
