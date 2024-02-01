import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import {environment} from "src/environments/environment.development"

@Component({
  selector: 'app-cookie-banner',
  templateUrl: './cookie-banner.component.html',
  styleUrls: ['./cookie-banner.component.css']
})
export class CookieBannerComponent {
  environment = environment
  ngOnInit(): void {
    
  }

  constructor(public dialog: MatDialog) {}
}
