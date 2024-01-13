import { Injectable } from '@angular/core';
import { ShoppingListEndpointService } from './shopping-list.endpoint.service';
import { ShoppingListItems } from './shopping-list.interface';
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

  constructor(
    private shoppingListEndpointService: ShoppingListEndpointService
  ) { }
}
