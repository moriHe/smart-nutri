import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Mealplan, PostMealplanPayload } from './mealplans.interface';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MealplansEndpointsService {
  fetchMealplan(date: string) {
    return this.http.get<Response<Mealplan>>(`http://localhost:8080/mealplan/${date}`)
  }

  postMealplanItem(payload: PostMealplanPayload) {
    return this.http.post<Response<string>>("http://localhost:8080/mealplan", payload)
  }

  constructor(
    private http: HttpClient
  ) {}
}
