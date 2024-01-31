import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-ingredient-source',
  templateUrl: './ingredient-source.component.html',
  styleUrls: ['./ingredient-source.component.css']
})
export class IngredientSourceComponent {
  @Input() url!: string

  openSource(event: Event) {
    event.stopPropagation()
    window.open(this.url, "_blank")
  }
}
