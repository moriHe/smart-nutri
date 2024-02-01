import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AddToShoppingList, ShoppingListCommonProps, ShoppingListByCategory } from './shopping-list.interface';
import { Response } from 'api';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListEndpointService {
  fetchShoppingList() {
    return this.http.get<Response<ShoppingListByCategory>>(`${environment.backendBaseUrl}/shopping-list`);
  }
  postShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/shopping-list/${mealplanId}`, shoppingListItems)
  }
  
  deleteShoppingListItem(id: number) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/mealplan/shopping-list/${id}`)
  }

  newDeleteShoppingListItems(ids: number[]) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/shopping-list/items/${ids.join(",")}`)
  }

  constructor(
    private http: HttpClient
  ) { }
}
