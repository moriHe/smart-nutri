import { Injectable } from '@angular/core';
import { Units } from 'api/recipes/recipes.interface';

@Injectable({
  providedIn: 'root'
})
export class UnitsService {

  UnitDisplay = {
    [Units.GRAM]: "g",
    [Units.MILLILITER]: "ml",
    [Units.TABLESPOON]: "EL",
    [Units.TEASPOON]: "TL",
  }

  ShoppingListUnitDisplay = {...this.UnitDisplay, PARTIAL: ""}

  constructor() { }
}
