import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent {
  openImprint() {
    this.router.navigate(["/impressum"])
  }
  openDataProtection() {
    this.router.navigate(["/datenschutz"])
  }

  constructor(private router: Router) {}
}
