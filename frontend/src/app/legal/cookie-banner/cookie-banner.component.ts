import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-cookie-banner',
  templateUrl: './cookie-banner.component.html',
  styleUrls: ['./cookie-banner.component.css']
})
export class CookieBannerComponent {

  ngOnInit(): void {
    
  }

  constructor(public dialog: MatDialog) {}
}
