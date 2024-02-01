import { Injectable } from '@angular/core';
import { Response } from 'api';
import { FullMealplanItem, Mealplans, PostMealplanPayload } from './mealplans.interface';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment.development';

export type FetchMealPlanQuery = {date: string, forShoppingList: boolean}

@Injectable({
  providedIn: 'root'
})
export class MealplansEndpointsService {
  fetchMealplan({date, forShoppingList}: FetchMealPlanQuery) {
    return this.http.get<Response<Mealplans>>(`${environment.backendBaseUrl}/mealplan/${date}?forShoppingList=${forShoppingList}`)
  }

  fetchMealplanItem(id: number) {
    return this.http.get<Response<FullMealplanItem>>(`${environment.backendBaseUrl}/mealplan/item/${id}`)
  }

  postMealplanItem(payload: PostMealplanPayload) {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/mealplan`, payload)
  }

  constructor(
    private http: HttpClient
  ) {}
}
