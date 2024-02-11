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
    [Units.PINCH]: "Prise",
    [Units.PIECE]: "Stk.",
    [Units.WHOLE]: "Ganze/r"
  }

  ShoppingListUnitDisplay = {...this.UnitDisplay, PARTIAL: ""}

  constructor() { }
}
