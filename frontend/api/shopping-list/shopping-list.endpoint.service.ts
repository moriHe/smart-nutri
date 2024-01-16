import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AddToShoppingList, ShoppingListItems } from './shopping-list.interface';
import { Response } from 'api';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListEndpointService {
  fetchShoppingList() {
    return this.http.get<Response<ShoppingListItems>>(`http://localhost:8080/mealplan/shopping-list`);
  }

  postShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.http.post<Response<string>>(`http://localhost:8080/shopping-list/${mealplanId}`, shoppingListItems)
  }

  constructor(
    private http: HttpClient
  ) { }
}
