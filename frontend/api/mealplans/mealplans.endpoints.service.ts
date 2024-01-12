import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Mealplan, PostMealplanPayload } from './mealplans.interface';
import { HttpClient } from '@angular/common/http';

export type FetchMealPlanQuery = {date: string, forShoppingList: boolean}

@Injectable({
  providedIn: 'root'
})
export class MealplansEndpointsService {
  fetchMealplan({date, forShoppingList}: FetchMealPlanQuery) {
    return this.http.get<Response<Mealplan>>(`http://localhost:8080/mealplan/${date}?forShoppingList=${forShoppingList}`)
  }

  postMealplanItem(payload: PostMealplanPayload) {
    return this.http.post<Response<string>>("http://localhost:8080/mealplan", payload)
  }

  constructor(
    private http: HttpClient
  ) {}
}
