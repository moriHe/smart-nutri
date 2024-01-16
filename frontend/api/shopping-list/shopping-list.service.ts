import { Injectable } from '@angular/core';
import { ShoppingListEndpointService } from './shopping-list.endpoint.service';
import { AddToShoppingList, ShoppingListItems } from './shopping-list.interface';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListService {

  getShoppingList() {
    return this.shoppingListEndpointService.fetchShoppingList().pipe(
      map((response: {data: ShoppingListItems}) => response ? response.data : [])
    )
  }

  addShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.shoppingListEndpointService.postShoppingListItems(shoppingListItems, mealplanId).pipe(
      map((response: {data: string}) => response)
    )
  }

  constructor(
    private shoppingListEndpointService: ShoppingListEndpointService
  ) { }
}
