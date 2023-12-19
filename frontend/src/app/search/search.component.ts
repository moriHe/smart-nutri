import { Component } from '@angular/core';
import { Result, TypesenseService } from '../typesense.service';
import { SearchResponseHit } from 'typesense/lib/Typesense/Documents';
import { RecipesService } from 'api/recipes/recipes.service';
import { ActivatedRoute } from '@angular/router';
import { Markets, Units } from 'api/recipes/recipes.interface';
import { FormBuilder, FormGroup } from '@angular/forms';
import { debounceTime, switchMap } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { SearchIngredientDialogComponent } from '../search-ingredient-dialog/search-ingredient-dialog.component';





@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  recipeId?: number
  results: SearchResponseHit<Result>[] = []

  ingredientInput!: FormGroup

  amountPerPortion: number = 1
  isBio: boolean = false
  selectedMarket: Markets = Markets.NONE
  markets: Markets[] = Object.values(Markets)
  selectedUnit: Units = Units.GRAM
  units: Units[] = Object.values(Units)

  ngOnInit(): void {
    this.route.params.subscribe(params => {
     this.recipeId = params['recipeId']
    })

    this.ingredientInput = this.formBuilder.group({
      query: ""
    })

    this.typesenseService.search("*").then((res) => {
      if (res) {
        this.results = res
      }
     })

    this.ingredientInput.get("query")?.valueChanges.pipe(
      debounceTime(800),
      switchMap((query: string) => this.typesenseService.search(query))
    ).subscribe((res) => {
      if (res) {
        return this.results = res
      }
      return this.results = []
    })
    
  
  }

  addIngredient(ingredientId: number) {
    if (this.recipeId) {
      this.recipesService.addRecipeIngredient(this.recipeId, {
        ingredientId,
        amountPerPortion: 3,
        isBio: false,
        market: Markets.REWE,
        unit: Units.GRAM
      }).subscribe()
    }
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(SearchIngredientDialogComponent, {
      data: {
        amountPerPortion: this.amountPerPortion,
         isBio: this.isBio, 
         selectedMarket: this.selectedMarket, 
         markets: this.markets, 
         selectedUnit: this.selectedUnit,
         units: this.units
        },
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result)
    });
  }

  constructor(
    private route: ActivatedRoute,
    private typesenseService: TypesenseService,
    private recipesService: RecipesService,
    private formBuilder: FormBuilder,
    public dialog: MatDialog
    ) { }

}
