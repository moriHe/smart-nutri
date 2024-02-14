import { ChangeDetectorRef, Component } from '@angular/core';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';
import { SnackbarService } from 'services/snackbar.service';


@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {
  hasRegistered = false

  email: string = '';
  password: string = '';

  singupErrorMessage = (errorMsg: string) => {
    let snackbarMsg = ""
    switch (errorMsg) {
      case "Unable to validate email address: invalid format":
        snackbarMsg = "Ungültiges E-Mail Format."
        break
      case "Password should be at least 6 characters.":
        snackbarMsg = "Das Passwort benötigt mindestens 6 Zeichen."
        break
      case "Email rate limit exceeded":
        snackbarMsg = "Wir können gerade keine weiteren Personen aufnehmen. Probieren Sie es später noch einmal."
        break
      default:
        snackbarMsg = "Etwas ging schief. Probieren Sie es später noch einmal."
    }
    this.snackbarService.openSnackbar(snackbarMsg, "Ok")
  }

  async signUp() {
    const response = await this.supabaseService.signUp(this.email, this.password)
    if (response.error) {
      console.log(response.error.message)
      this.singupErrorMessage(response.error.message)
    } else {
      this.hasRegistered = true

    }
  }

  goToLogin() {
    this.router.navigate(["/"])
  }

  constructor(
    private snackbarService: SnackbarService,
    private router: Router,
    private supabaseService: SupabaseService
    ) { }
}
