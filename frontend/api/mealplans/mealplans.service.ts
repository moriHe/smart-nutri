import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import { FullMealplanItem, Mealplans, PostMealplanPayload } from './mealplans.interface';
import { Response } from 'api';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment.development'

type FetchMealPlanQuery = {date: string, forShoppingList: boolean}
@Injectable({
  providedIn: 'root'
})

export class MealplansService {
  getMealplan(query: FetchMealPlanQuery): Observable<Mealplans> {
    return this.http.get<Response<Mealplans>>(`${environment.backendBaseUrl}/mealplan/${query.date}?forShoppingList=${query.forShoppingList}`)
    .pipe(
      map((response: { data: Mealplans }) => response ? response.data : [])
    )
  }

  getMealplanItem(id: number): Observable<FullMealplanItem> {
    return this.http.get<Response<FullMealplanItem>>(`${environment.backendBaseUrl}/mealplan/item/${id}`).pipe(
      map((response: {data: FullMealplanItem}) => response.data)
    )
  }

  addMealplanItem(payload: PostMealplanPayload): Observable<string> {
    return this.http.post<Response<string>>(`${environment.backendBaseUrl}/mealplan`, payload).pipe(
      map((response: {data: string}) => response.data)
    )
  }

  constructor(private http: HttpClient) { }
}
