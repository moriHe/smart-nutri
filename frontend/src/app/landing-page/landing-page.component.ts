import { Component } from '@angular/core';
import { signInWithEmailAndPassword } from "firebase/auth";
import {Auth} from '@angular/fire/auth'
import { FormBuilder } from '@angular/forms';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent {
  loginInput = this.formBuilder.group({
    email: "",
    password: ""
  })

  onLogin() {
    if (this.loginInput.value.email && this.loginInput.value.password) {
      signInWithEmailAndPassword(this.auth, this.loginInput.value.email, this.loginInput.value.password).then((result) => {
      })
    }
  }



    constructor(
      private auth: Auth, 
      private formBuilder: FormBuilder,
      ) {}
  }
