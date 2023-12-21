import { Component, Inject } from "@angular/core";
import {
    MAT_DIALOG_DATA,
    MatDialogRef,
  } from '@angular/material/dialog';
import { Markets, Units } from "api/recipes/recipes.interface";
import { MarketsService } from "services/markets.service";
import { UnitsService } from "services/units.service";

export interface DialogData {
  amountPerPortion: number,
  isBio: boolean,
  selectedMarket: Markets,
  markets: Markets[],
  selectedUnit: Units,
  units: Units[],
}

@Component({
  selector: 'app-search-ingredient-dialog',
  templateUrl: './search-ingredient-dialog.component.html',
  styleUrls: ['./search-ingredient-dialog.component.css']
})
export class SearchIngredientDialogComponent {
  onNoClick(): void {
    this.dialogRef.close();
  }

  toggleIsBio() {
    this.data.isBio = !this.data.isBio
  }

  constructor(
    public dialogRef: MatDialogRef<SearchIngredientDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}

}
