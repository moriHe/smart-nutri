import { Component } from '@angular/core';
import { Result, TypesenseService } from 'api/ingredient-search/typesense.service';
import { SearchResponseHit } from 'typesense/lib/Typesense/Documents';
import { RecipesService } from 'api/recipes/recipes.service';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup } from '@angular/forms';
import { debounceTime, switchMap, take } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { SearchIngredientDialogComponent } from '../search-ingredient-dialog/search-ingredient-dialog.component';
import { MatSnackBar, MatSnackBarRef, SimpleSnackBar } from '@angular/material/snack-bar';
import { FullRecipe } from 'api/recipes/recipes.interface';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  recipe!: FullRecipe
  results: SearchResponseHit<Result>[] = []

  ingredientInput!: FormGroup

  ingredientId!: number
  ingredientName!: string

  ngOnInit(): void {
    this.route.params.subscribe(params => {
     this.recipesService.getRecipe(params['recipeId']).pipe(take(1)).subscribe((response: FullRecipe) => {
      this.recipe = response
    });
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

  openDialog(ingredientId: number, ingredientName: string): void {
    this.ingredientId = ingredientId
    this.ingredientName = ingredientName
    const dialogRef = this.dialog.open(SearchIngredientDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (this.recipe.id && result) {
        this.recipesService.addRecipeIngredient(this.recipe.id, {
          ingredientId: this.ingredientId,
          amountPerPortion: Number(result.amountPerPortion) / this.recipe.defaultPortions,
          isBio: result.isBio,
          market: result.selectedMarket,
          unit: result.selectedUnit
        }).subscribe({
          next: (response) => {
            this.openSnackbar({type: "SUCCESS", recipeIngredientId: response})

          },
          error: () => {
            this.openSnackbar({type: "ERROR"})
          }
        })
      }
    });
  }

  openSnackbar({type, recipeIngredientId = 0}: {type: "SUCCESS" | "ERROR", recipeIngredientId?: number}) {
    if (type === "SUCCESS") {
      const snackBarRef: MatSnackBarRef<SimpleSnackBar> = this.snackbar.open(
        `Hinzugefügt: ${this.ingredientName}`,
         "Rückgängig", 
         {
        horizontalPosition: "start",
        verticalPosition: "bottom",
        duration: 3000
        })
        snackBarRef.onAction().subscribe(() => {
          this.recipesService.removeRecipeIngredient(recipeIngredientId).subscribe()
        });
    } else {
      this.snackbar.open("Etwas ging schief.", "Ok", {
        horizontalPosition: "start",
        verticalPosition: "bottom"
    })
  }
}

  constructor(
    private route: ActivatedRoute,
    private typesenseService: TypesenseService,
    private recipesService: RecipesService,
    private formBuilder: FormBuilder,
    public dialog: MatDialog,
    private snackbar: MatSnackBar,
    ) { }

}
