import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { Mealplan } from './mealplans.interface';
import { MealplansEndpointsService } from './mealplans.endpoints.service';

@Injectable({
  providedIn: 'root'
})
export class MealplansService {
  getMealplan(date: string): Observable<Mealplan> {
    return this.mealplanEndpoint.fetchMealplan(date).pipe(
      map((response: { data: Mealplan }) => response ? response.data : [])
    )
  }

  constructor(private mealplanEndpoint: MealplansEndpointsService) { }
}
