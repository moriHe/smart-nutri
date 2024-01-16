import { Injectable } from '@angular/core';
import { Response } from 'api';
import { FullMealplanItem, Mealplans, PostMealplanPayload } from './mealplans.interface';
import { HttpClient } from '@angular/common/http';

export type FetchMealPlanQuery = {date: string, forShoppingList: boolean}

@Injectable({
  providedIn: 'root'
})
export class MealplansEndpointsService {
  fetchMealplan({date, forShoppingList}: FetchMealPlanQuery) {
    return this.http.get<Response<Mealplans>>(`http://localhost:8080/mealplan/${date}?forShoppingList=${forShoppingList}`)
  }

  fetchMealplanItem(id: number) {
    return this.http.get<Response<FullMealplanItem>>(`http://localhost:8080/mealplan/item/${id}`)
  }

  postMealplanItem(payload: PostMealplanPayload) {
    return this.http.post<Response<string>>("http://localhost:8080/mealplan", payload)
  }

  constructor(
    private http: HttpClient
  ) {}
}
