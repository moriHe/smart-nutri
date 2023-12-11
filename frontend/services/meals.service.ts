import { Injectable } from '@angular/core';
import { Meals } from 'api/recipes/recipes.interface';

@Injectable({
  providedIn: 'root'
})
export class MealsService {

  MealDisplay = {
    [Meals.NONE]: "-",
    [Meals.BREAKFAST]: "Frühstück",
    [Meals.LUNCH]: "Mittagessen",
    [Meals.DINNER]: "Abendessen"
}

  constructor() { }
}
