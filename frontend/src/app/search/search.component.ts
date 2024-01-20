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
import algoliasearch from 'algoliasearch/lite';
import { environment } from 'src/environments/environment.development';

type AlgoliaResult = {
  objectID: string
  id: number
  name: string
  brands: string
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  currentIngredientName!: string
  recipe!: FullRecipe
  results: SearchResponseHit<Result>[] = []

  searchClient = algoliasearch(environment.algoliaAppId, environment.algoliaSearchOnlyApiKey);
  index = this.searchClient.initIndex(environment.indexName);

  query: string = '';
  algoliaResults: AlgoliaResult[] = [];

  onSearch(): void {
    if (this.query.length > 0) {
      this.index.search<AlgoliaResult>(this.query).then(({ hits }) => {
        this.algoliaResults = hits
      });
    } else {
      this.results = [];
    }
  }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
     this.recipesService.getRecipe(params['recipeId']).pipe(take(1)).subscribe((response: FullRecipe) => {
      this.recipe = response
    });
    })
  }

  openDialog(ingredientId: number, ingredientName: string): void {
    this.currentIngredientName = ingredientName
    const dialogRef = this.dialog.open(SearchIngredientDialogComponent);

    dialogRef.afterClosed().subscribe(result => {
      if (this.recipe.id && result) {
        this.recipesService.addRecipeIngredient(this.recipe.id, {
          ingredientId,
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
        `Hinzugefügt: ${this.currentIngredientName}`,
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
    private recipesService: RecipesService,
    private formBuilder: FormBuilder,
    public dialog: MatDialog,
    private snackbar: MatSnackBar,
    ) { }

}
