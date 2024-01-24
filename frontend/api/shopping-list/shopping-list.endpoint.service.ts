import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AddToShoppingList, ShoppingListCommonProps, ShoppingListByCategory } from './shopping-list.interface';
import { Response } from 'api';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListEndpointService {
  fetchShoppingList() {
    return this.http.get<Response<ShoppingListByCategory>>(`http://localhost:8080/shopping-list`);
  }
  postShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.http.post<Response<string>>(`http://localhost:8080/shopping-list/${mealplanId}`, shoppingListItems)
  }
  
  deleteShoppingListItem(id: number) {
    return this.http.delete<Response<string>>(`http://localhost:8080/mealplan/shopping-list/${id}`)
  }

  constructor(
    private http: HttpClient
  ) { }
}
