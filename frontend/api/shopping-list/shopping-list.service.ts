import { Injectable } from '@angular/core';
import { AddToShoppingList, ShoppingListByCategory } from './shopping-list.interface';
import { map } from 'rxjs';
import { Response } from 'api';
import { environment } from 'src/environments/environment.development'
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ShoppingListService {

  getShoppingList() {
    return this.http.get<Response<ShoppingListByCategory>>(`${environment.backendBaseUrl}/shopping-list`)
      .pipe(
        map((response: {data: ShoppingListByCategory}) => response?.data)
      )
  }

  addShoppingListItems(shoppingListItems: AddToShoppingList[], mealplanId: number) {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/shopping-list/${mealplanId}`, shoppingListItems)
      .pipe(
        map((response: {data: string}) => response)
      )
  }

  removeShoppingListItem(id: number) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/mealplan/shopping-list/${id}`)
      .pipe(
        map((
      response: {data: string}) => response)
      )
  }

  newRemoveShoppingListItems(ids: number[]) {
    return this.http.delete<Response<string>>(`${environment.backendBaseUrl}/shopping-list/items/${ids.join(",")}`)
    .pipe(
      map((response: {data: string}) => response)
    )
  }

  constructor(
    private http: HttpClient
  ) { }
}
