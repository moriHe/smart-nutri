import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { SupabaseService } from 'api/supabase.service';
import { Router } from '@angular/router';
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

  async onLogin() {
    if (this.loginInput.value.email && this.loginInput.value.password) {
      const response = await this.supabaseService.login(this.loginInput.value.email, this.loginInput.value.password)
      if (response.data.session) {
        this.supabaseService.setSession(response.data.session)
        this.router.navigate(['/meine-rezepte']);
      }
    }
    }
  



    constructor(
      private router: Router,
      private supabaseService: SupabaseService,
      private formBuilder: FormBuilder,
      ) {}
  }
