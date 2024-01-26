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
    const response = await this.supabase.auth.signUp({email, password, options: {emailRedirectTo: "http://localhost:4200/willkommen"}})
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

  async initialize() {
    const cookieConsent = localStorage.getItem("cookieConsent")
    if (cookieConsent === null) {
      return
    }

    try {
      const initialSession = await this.supabase.auth.getSession();
      this.setSession(initialSession.data.session);
      this.supabase.auth.onAuthStateChange((_, session) => {
        this.sessionSubject.next(session);
      });
    await firstValueFrom(this.userService.getUser())
  } catch (error) {
    console.log(error)
  }
    this.isAppInitializedSubject.next(true)
  
  }
  
  constructor(private userService: UserService) {}
}
