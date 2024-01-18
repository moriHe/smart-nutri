import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-shopping-list-bell',
  templateUrl: './shopping-list-bell.component.html',
  styleUrls: ['./shopping-list-bell.component.css']
})
export class ShoppingListBellComponent {
  @Input() count = 0
  @Input() onClick!: () => void | undefined
  
  onClickBell() {
    if (this.onClick) {
      this.onClick()
    }
  }
}
