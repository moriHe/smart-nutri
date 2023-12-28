import { Component } from "@angular/core";
import { MatDialogRef } from '@angular/material/dialog';
import { Markets, Units } from "api/recipes/recipes.interface";
import { MarketsService } from "services/markets.service";
import { UnitsService } from "services/units.service";
@Component({
  selector: 'app-search-ingredient-dialog',
  templateUrl: './search-ingredient-dialog.component.html',
  styleUrls: ['./search-ingredient-dialog.component.css']
})
export class SearchIngredientDialogComponent {
  amountPerPortion: number = 1
  isBio: boolean = false
  selectedMarket: Markets = Markets.NONE
  markets: Markets[] = Object.values(Markets)
  selectedUnit: Units = Units.GRAM
  units: Units[] = Object.values(Units)

  onNoClick(): void {
    this.dialogRef.close();
  }

  toggleIsBio() {
    this.isBio = !this.isBio
  }

  constructor(
    public dialogRef: MatDialogRef<SearchIngredientDialogComponent>,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}

}
