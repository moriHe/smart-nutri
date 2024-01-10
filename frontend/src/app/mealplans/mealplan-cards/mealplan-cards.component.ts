import { Component, Input } from '@angular/core';
import { Mealplan } from 'api/mealplans/mealplans.interface';
import { Meals } from 'api/recipes/recipes.interface';
import { MealsService } from 'services/meals.service';

@Component({
  selector: 'app-mealplan-cards',
  templateUrl: './mealplan-cards.component.html',
  styleUrls: ['./mealplan-cards.component.css']
})
export class MealplanCardsComponent {
  @Input() mealKey!: Meals;
  @Input() mealplan: Mealplan = []
  @Input() openDialog!: (mealKey: Meals) => void

  
  getMealplanForMealType(): Mealplan {
    return this.mealplan.filter(item => item.meal === Meals[this.mealKey]);
  }

  constructor(public mealsService: MealsService) { }
}
