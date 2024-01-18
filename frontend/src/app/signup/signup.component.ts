import { Component } from '@angular/core';
import { Auth } from '@angular/fire/auth';
import { Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';
import firebase from 'firebase/compat/app';
import 'firebase/compat/auth';
import * as firebaseui from "firebaseui"

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

  ngOnInit(): void {
    const ui = new firebaseui.auth.AuthUI(this.auth)

    ui.start('#firebaseui-auth-container', {
      callbacks: {
        signInSuccessWithAuthResult: (result) => {
          this.userService.addUser().subscribe(() => {
            this.router.navigate(['/willkommen']);
          })
          return false
        }
      },
      signInOptions: [
        firebase.auth.EmailAuthProvider.PROVIDER_ID,
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
      ],
    });
  }

  constructor(
    private supabaseService: SupabaseService,
    private router: Router,
    private auth: Auth,
    private userService: UserService
    ) { }
}
