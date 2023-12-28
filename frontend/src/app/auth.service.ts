import { Injectable } from '@angular/core';
import { Auth, authState, User } from '@angular/fire/auth';
import { Subscription } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  authState$ = authState(this.auth);
  authStateSubscription!: Subscription;
  currentUser!: null | User

  getCurrentUser(): null | User {
    return this.currentUser
  }

  
  constructor(private auth: Auth) { }
}
