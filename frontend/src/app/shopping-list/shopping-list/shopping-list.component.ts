import { Component } from '@angular/core';
import { Mealplans } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { take } from 'rxjs';

@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  mealplansNotOnShoppingList: Mealplans = []
  // todo get mealplans that are not on the shopping list
  // from today to future

// todo get shopping list items (dont forget new column isInShoppingList on mealplan)
  // from today to future

ngOnInit(): void {
  this.mealplansService.getMealplan({date: new Date().toISOString(), forShoppingList: true})
    .pipe(take(1)).subscribe((response) => {
      if (response) {
        this.mealplansNotOnShoppingList = response
      }
    })
}
  constructor(
    private mealplansService: MealplansService
  ) {}
}
