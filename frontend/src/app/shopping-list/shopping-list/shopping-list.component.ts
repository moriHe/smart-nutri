import { Component } from '@angular/core';
import { MealplansService } from 'api/mealplans/mealplans.service';

@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  // todo get mealplans that are not on the shopping list
  // from today to future
  
// todo get shopping list items (dont forget new column isInShoppingList on mealplan)
  // from today to future


  constructor(
    mealplansService: MealplansService
  ) {}
}
