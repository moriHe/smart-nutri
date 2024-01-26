import { ChangeDetectorRef, Component } from '@angular/core';
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
    console.log(isSuccess)
    if (isSuccess) {
      this.hasRegistered = true
    }
  }

  constructor(
    private supabaseService: SupabaseService
    ) { }
}
