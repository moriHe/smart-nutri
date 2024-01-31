import { Component, ChangeDetectorRef} from '@angular/core';
import { UserService } from 'api/user/user.service';
import { SupabaseService } from 'api/supabase.service';
import { finalize } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { CookieBannerComponent } from './legal/cookie-banner/cookie-banner.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Smart Nutri';
  // isInitialized = this.supabaseService.isAppIntialized$
    isValidSecret!: boolean
 



  ngOnInit(): void {
    this.supabaseService.isValidSecret$
      .subscribe((value) => {
        this.isValidSecret = value;
      });

    const cookieConsent = localStorage.getItem("cookieConsent")
    const currentUrl = window.location.href
    const legalPages = currentUrl.endsWith("/datenschutz") || currentUrl.endsWith("/impressum") || currentUrl.endsWith("/datenbank-nahrungsmittel")
    if (legalPages) {
      return
    }
    if (cookieConsent === null) {
      const dialogRef = this.dialog.open(CookieBannerComponent, {disableClose: true, width: "50%", enterAnimationDuration: "1000ms", exitAnimationDuration: "1000ms"})
      dialogRef.afterClosed().subscribe((result) => {
        if (result.cookieConsent) {
          localStorage.setItem("cookieConsent", "true")
        }
      })
    }
  }

  constructor(
    public dialog: MatDialog,
    private supabaseService: SupabaseService
    ) {}
}
