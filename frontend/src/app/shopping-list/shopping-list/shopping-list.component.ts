import { ChangeDetectorRef, Component } from '@angular/core';
import { MarketsService } from 'services/markets.service';
import { UnitsService } from 'services/units.service';


@Component({
  selector: 'app-shopping-list',
  templateUrl: './shopping-list.component.html',
  styleUrls: ['./shopping-list.component.css']
})
export class ShoppingListComponent {
  isShoppingListView = true

openAddItemsView = () => {
  this.isShoppingListView = false
  this.cdr.detectChanges()
}

openShoppingListView = () => {
  this.isShoppingListView = true
  this.cdr.detectChanges()
}

  constructor(
    private cdr: ChangeDetectorRef,
    public marketsService: MarketsService,
    public unitsService: UnitsService
  ) {}
}
