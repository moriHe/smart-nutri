import { Injectable } from '@angular/core';
import { Response } from 'api';
import { Mealplan } from './mealplans.interface';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MealplansEndpointsService {
  fetchMealplan(date: string) {
    return this.http.get<Response<Mealplan>>(`http://localhost:8080/mealplan/${date}`)
  }

  // RecipeId int     `json:"recipeId"`
	// Date     string  `json:"date"`
	// Meal     string  `json:"meal"`
	// Portions float32 `json:"portions"`

  postMealplanItem() {
    return this.http.post<Response<string>>("http://localhost:8080/mealplan", {
      recipeId: 1,
      date: new Date().toISOString(),
      meal: 'BREAKFAST',
      portions: 3
    })
  }

  constructor(
    private http: HttpClient
  ) {}
}
