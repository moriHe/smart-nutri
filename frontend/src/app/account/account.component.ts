import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent {

  deleteAccount() {
    this.userService.deleteUser().subscribe(() => {
      this.supabaseService.sessionSubject.next(null)
      this.userService.setUserNull()
      this.supabaseService.logout()
      this.router.navigate(["/"])
    })
  }

  constructor(
    private supabaseService: SupabaseService,
    private router: Router,
    private userService: UserService
    ) {}
}
