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
  email: string = '';
  password: string = '';

  async onLogin() {
      const response = await this.supabaseService.login(this.email, this.password)
      if (response.data.session) {
        this.supabaseService.setSession(response.data.session)
        this.router.navigate(['/meine-rezepte']);
      }
    
    }
  



    constructor(
      private router: Router,
      private supabaseService: SupabaseService,
      private formBuilder: FormBuilder,
      ) {}
  }
