import { Component, OnInit } from '@angular/core';
import { Recipe, RecipesService } from '../recipes.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  constructor(private recipesService: RecipesService) { }

  recipes: Recipe[] = []

  ngOnInit(): void {
    this.recipesService.getRecipes().subscribe((response: Recipe[]) => {
      this.recipes = response
    })
    
  }


}
