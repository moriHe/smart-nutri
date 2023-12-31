import { Component } from '@angular/core';
import { signInWithEmailAndPassword } from "firebase/auth";
import {Auth} from '@angular/fire/auth'
import { Router } from '@angular/router';
import { FormBuilder, FormGroup } from '@angular/forms';
import { UserService } from 'api/user/user.service';
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
        console.log(result)
      })
    }
  }



    constructor(
      private auth: Auth, 
      private router: Router,
      private formBuilder: FormBuilder,
      ) {}
  }
