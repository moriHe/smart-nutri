import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { SupabaseService } from 'api/supabase.service';
import { Router } from '@angular/router';
import { UserService } from 'api/user/user.service';
import { SnackbarService } from 'services/snackbar.service';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent {
  email: string = '';
  password: string = '';

  async onLogin() {
      const response = await this.supabaseService.login(this.email, this.password)
      if (response.data.session) {
        await this.supabaseService.initialize()
        this.router.navigate(['/meine-rezepte']);        
      } else if (response.error) {
        let snackbarMsg = "Etwas ging schief."
        if (response.error.message === "Invalid login credentials") {
          snackbarMsg = "Ungültige Logindaten"
        }
        this.snackbarService.openSnackbar(snackbarMsg, "Ok")
      }
    
    }
  
    goToRegister() {
      this.router.navigate(["/registrieren"])
    }



    constructor(
      private snackbarService: SnackbarService,
      private router: Router,
      private supabaseService: SupabaseService
      ) {}
  }
