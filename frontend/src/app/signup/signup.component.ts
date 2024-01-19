import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {
  hasRegistered = false

  email: string = '';
  password: string = '';

  async signUp() {
    const isSuccess = await this.supabaseService.signUp(this.email, this.password)
    if (isSuccess) {
      this.userService.addUser().subscribe(() => {
        this.router.navigate(['/willkommen']);
      })
    }
  }
  constructor(
    private supabaseService: SupabaseService,
    private router: Router,
    private userService: UserService
    ) { }
}
