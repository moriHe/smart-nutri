import { Component } from '@angular/core';
import { RecipesService } from 'api/recipes/recipes.service';
import { ActivatedRoute } from '@angular/router';
import { Subject, debounceTime, take } from 'rxjs';
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
  url: string
  code: string
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  private searchSubject = new Subject<string>();

  currentIngredientName!: string
  recipe!: FullRecipe

  searchClient = algoliasearch(environment.algoliaAppId, environment.algoliaSearchOnlyApiKey);
  index = this.searchClient.initIndex(environment.indexName);

  query: string = '';
  algoliaResults: AlgoliaResult[] = [];

  onSearch(): void {
    this.searchSubject.next(this.query);
  }

  private performSearch(query: string, initial = false): void {
    if (initial || query.length > 0) {
      this.index.search<AlgoliaResult>(query).then(({ hits }) => {
        this.algoliaResults = hits;
      });
    } else {
      this.algoliaResults = [];
    }
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

openSource(url: string, event: Event) {
  event.stopPropagation()
  window.open(url, "_blank")
}

ngOnInit(): void {
  this.route.params.subscribe(params => {
   this.recipesService.getRecipe(params['recipeId']).pipe(take(1)).subscribe((response: FullRecipe) => {
    this.recipe = response
  });
  })

  this.performSearch("", true)
  this.searchSubject.pipe(debounceTime(500)).subscribe((query) => {
    this.performSearch(query);
  });
}

  constructor(
    private route: ActivatedRoute,
    private recipesService: RecipesService,
    public dialog: MatDialog,
    private snackbar: MatSnackBar,
    ) { }

}
