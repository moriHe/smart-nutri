import { Component } from '@angular/core';
import { Mealplans } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { ShoppingListItems } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { take } from 'rxjs';

@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  mealplansNotOnShoppingList: Mealplans = []
  shoppingList: ShoppingListItems = []
  
  // TODO: Add mealplan items to shopping list
  // *** kind of a session approach (abbrechen, speichern und weiter, speichern und beenden)
  // display shopping list items
  // remove shopping list items

ngOnInit(): void {
  this.mealplansService.getMealplan({date: new Date().toISOString(), forShoppingList: true})
    .pipe(take(1)).subscribe((response) => {
      if (response) {
        this.mealplansNotOnShoppingList = response
      }
    })

    this.shoppingListService.getShoppingList().pipe(take(1)).subscribe((response) => {
      console.log("test")
      if (response) {
        this.shoppingList = response
      }
    })
}
  constructor(
    private shoppingListService: ShoppingListService,
    private mealplansService: MealplansService
  ) {}
}
