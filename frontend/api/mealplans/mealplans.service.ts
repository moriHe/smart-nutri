import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { Mealplans, PostMealplanPayload } from './mealplans.interface';
import { FetchMealPlanQuery, MealplansEndpointsService } from './mealplans.endpoints.service';

@Injectable({
  providedIn: 'root'
})
export class MealplansService {
  getMealplan(query: FetchMealPlanQuery): Observable<Mealplans> {
    return this.mealplanEndpoint.fetchMealplan(query).pipe(
      map((response: { data: Mealplans }) => response ? response.data : [])
    )
  }

  addMealplanItem(payload: PostMealplanPayload): Observable<string> {
    return this.mealplanEndpoint.postMealplanItem(payload).pipe(
      map((response: {data: string}) => response.data)
    )
  }

  constructor(private mealplanEndpoint: MealplansEndpointsService) { }
}
