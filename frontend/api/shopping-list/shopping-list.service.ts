import { Injectable } from '@angular/core';
import { ShoppingListEndpointService } from './shopping-list.endpoint.service';
import { AddToShoppingList, ShoppingListByCategory } from './shopping-list.interface';
import { map } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListService {

  getShoppingList() {
    return this.shoppingListEndpointService.fetchShoppingList().pipe(
      map((response: {data: ShoppingListByCategory}) => response?.data)
    )
  }

  addShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.shoppingListEndpointService.postShoppingListItems(shoppingListItems, mealplanId).pipe(
      map((response: {data: string}) => response)
    )
  }

  removeShoppingListItem(id: number) {
    return this.shoppingListEndpointService.deleteShoppingListItem(id).pipe(map((
      response: {data: string}) => response))
  }

  newRemoveShoppingListItems(ids: number[]) {
    return this.shoppingListEndpointService.newDeleteShoppingListItems(ids).pipe(map((
      response: {data: string}) => response))
  }

  constructor(
    private shoppingListEndpointService: ShoppingListEndpointService
  ) { }
}
