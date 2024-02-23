import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';
import { SnackbarService } from 'services/snackbar.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent {

  deleteAccount() {
    this.userService.deleteUser().subscribe({
      next: () => {
        this.supabaseService.sessionSubject.next(null)
        this.userService.setUserNull()
        this.supabaseService.logout()
        this.router.navigate(["/"])
      },
      error: () => {
        this.snackbarService.openSnackbar("Etwas ging schief. Bitte melden Sie sich beim Support um den Account manuell zu löschen.", "Ok")
      }
    })
  }

  constructor(
    private snackbarService: SnackbarService,
    private supabaseService: SupabaseService,
    private router: Router,
    private userService: UserService
    ) {}
}
