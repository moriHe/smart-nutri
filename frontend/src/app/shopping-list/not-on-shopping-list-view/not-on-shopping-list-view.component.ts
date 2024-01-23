import { Component, Input } from '@angular/core';
import { FullMealplanItem } from 'api/mealplans/mealplans.interface';
import { MealplansService } from 'api/mealplans/mealplans.service';
import { AddToShoppingList, RecipeIngredientItem } from 'api/shopping-list/shopping-list.interface';
import { ShoppingListService } from 'api/shopping-list/shopping-list.service';
import { take } from 'rxjs';

@Component({
  selector: 'app-not-on-shopping-list-view',
  templateUrl: './not-on-shopping-list-view.component.html',
  styleUrls: ['./not-on-shopping-list-view.component.css']
})
export class NotOnShoppingListViewComponent {
  @Input() openShoppingListView!: () => void
  mealplanNumberNotOnShoppingList = 0
  mealplanItem: FullMealplanItem | undefined = undefined

  addMealplanToShoppingList(action: "finish" | "next") {
    let updatedAddToShoppingList: AddToShoppingList[] = []
    if (this.mealplanItem) {
      const shoppingListItems: AddToShoppingList[] = this.mealplanItem.recipe.recipeIngredients.map(
        (item) => {
          return {
            recipeIngredientId: item.id,
            market: item.market,
            isBio: item.isBio
          }
        }
      )
      updatedAddToShoppingList = shoppingListItems
    }

      this.shoppingListService.addShoppingListItems(updatedAddToShoppingList, this.mealplanItem!.id).pipe(take(1)).subscribe((response) => {
        if (response) {
         if (action === "finish") {
          this.openShoppingListView()
          return
         } else {
           this.updateListItems()
         }
        }
      })
  }


  removeItem(recipeIngredientId: number) {
    if (this.mealplanItem) {
      this.mealplanItem = {
        ...this.mealplanItem,
        recipe: {
          ...this.mealplanItem.recipe,
          recipeIngredients: this.mealplanItem.recipe.recipeIngredients.filter(
            (recipeIngredient: RecipeIngredientItem) => recipeIngredient.id !== recipeIngredientId)
        }
      }
    }
  }


  updateListItems() {
    this.mealplansService.getMealplan({date: new Date().toISOString(), forShoppingList: true})
      .pipe(take(1)).subscribe((response) => {
        if (response) {
          if (response.length > 0) {
            this.mealplansService.getMealplanItem(response[0].id).pipe(take(1))
              .subscribe((response) => {
              this.mealplanItem = response

              })
          }
          this.mealplanNumberNotOnShoppingList = response.length
        }
      })
  }
  
  ngOnInit(): void {
    this.updateListItems()

      
  }

  constructor(
    private mealplansService: MealplansService,
    private shoppingListService: ShoppingListService
  ) {}
}
