import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ShoppingListItems } from './shopping-list.interface';
import { Response } from 'api';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListEndpointService {
  fetchShoppingList() {
    return this.http.get<Response<ShoppingListItems>>(`http://localhost:8080/mealplan/shopping-list`);
  }
  constructor(
    private http: HttpClient
  ) { }
}
