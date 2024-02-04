import { Injectable } from '@angular/core';
import {
  AuthChangeEvent,
  AuthSession,
  createClient,
  Session,
  SupabaseClient,
  User,
} from '@supabase/supabase-js'
import { BehaviorSubject, firstValueFrom, Observable, Subject } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { UserService } from './user/user.service';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class SupabaseService {
  private supabase: SupabaseClient = createClient(environment.supabaseUrl, environment.supabaseKey)
  _session: AuthSession | null = null

  private isAppInitializedSubject: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)
  isAppIntialized$: Observable<boolean> = this.isAppInitializedSubject.asObservable()

  sessionSubject: BehaviorSubject<AuthSession | null> = new BehaviorSubject<AuthSession | null>(null);
  // Observable for session changes
  session$: Observable<AuthSession | null> = this.sessionSubject.asObservable();

  isValidSecretSubject: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)
  isValidSecret$: Observable<boolean> = this.isValidSecretSubject.asObservable();

  get session(): AuthSession | null {
    return this.sessionSubject.value;
  }

  // Setter for updating the session
  setSession(currentSession: AuthSession | null) {
    this.sessionSubject.next(currentSession);
  }

  authChanges(callback: (event: AuthChangeEvent, session: Session | null) => void) {
    return this.supabase.auth.onAuthStateChange(callback)
  }

  async signUp(email: string, password: string) {
    const response = await this.supabase.auth.signUp({email, password, options: {emailRedirectTo: `${environment.frontendBaseUrl}/willkommen/`}})
    if (!response.error) {
      return true
    }
    return false
  }

  async login(email: string, password: string) {
    return await this.supabase.auth.signInWithPassword({email, password})
  }

  async logout() {
    return await this.supabase.auth.signOut()
  }

  async initialize(): Promise<void> { 
    const secret = localStorage.getItem("secret")
    if (!secret) {
      const  enteredSecret = prompt("This is a portfolio page and not open to the public. Please enter the secret to access the site: ")
      if (enteredSecret) {
        localStorage.setItem("secret", enteredSecret)
      }
    }

      try {
        const response = await firstValueFrom(this.userService.getSecret())
        if (response) {
          this.isValidSecretSubject.next(true)
        }
      } catch(error) {
        this.isValidSecretSubject.next(false)
        this.sessionSubject.next(null)
        return
      }
    
      const cookieConsent = localStorage.getItem("cookieConsent");
    if (!cookieConsent) {
      return;
    }


    try {
      await new Promise<void>((resolve) => {
        this.supabase.auth.onAuthStateChange((_, session) => {
          this.sessionSubject.next(session);
            resolve();
        })
      });
  
      if (this.session) {
      await firstValueFrom(this.userService.getUser());
    }
    } catch (error) {
      if (this.router.url.includes("/willkommen")) {
        return
      }
      console.error(error);
    }
  
    this.isAppInitializedSubject.next(true);
  }
  
  constructor(private userService: UserService, private router: Router) {}
}
