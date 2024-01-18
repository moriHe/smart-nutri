import { Injectable } from '@angular/core';
import {
  AuthChangeEvent,
  AuthSession,
  createClient,
  Session,
  SupabaseClient,
  User,
} from '@supabase/supabase-js'
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class SupabaseService {
  private supabase: SupabaseClient = createClient(environment.supabaseUrl, environment.supabaseKey)
  _session: AuthSession | null = null

  private sessionSubject: BehaviorSubject<AuthSession | null> = new BehaviorSubject<AuthSession | null>(null);
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
    const response = await this.supabase.auth.signUp({email, password})
    console.log(response)
    if (!response.error) {
      return true
    }
    return false
  }

  async initialize() {
    // You can include additional initialization logic if needed
    // ...

    // For example, fetch the initial session
    const initialSession = await this.supabase.auth.getSession();
    this.setSession(initialSession.data.session);
  }

  constructor() {
    this.supabase.auth.onAuthStateChange((_, session) => {
      this.sessionSubject.next(session);
    });
  }
}
