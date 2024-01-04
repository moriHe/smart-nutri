import { Component } from '@angular/core';
import { Auth } from '@angular/fire/auth';
import { Router } from '@angular/router';
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
    private router: Router,
    private auth: Auth,
    private userService: UserService
    ) { }
}