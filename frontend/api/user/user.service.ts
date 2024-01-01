import { Injectable, signal } from '@angular/core';
import { Api, Response } from 'api';
import { BehaviorSubject, Observable, finalize, map, of, switchMap, take } from 'rxjs';
import { Auth, idToken, authState, User } from '@angular/fire/auth';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private isInitializedSubject = new BehaviorSubject<boolean>(false);
  isInitialized$ = this.isInitializedSubject.asObservable();
  
  authState$ = authState(this.auth);
 
  user = this.authState$.pipe(
    take(1),
    switchMap((authUser) => {
    if (authUser) {
      return this.api.fetchUser()
    }
    return of(null)
    }),
    finalize(() => this.isInitializedSubject.next(true)))
    .subscribe((user) => {
      console.log(user)
      return user
    })
  
  isAuthRegisteredUser() {
    return this.authState$.pipe((map(user => {
      if (user) {
        return true
      }
      return this.router.createUrlTree([""])
    })))
  }



  private userIdSubject = new BehaviorSubject<number | null>(null);
  userId$ = this.userIdSubject.asObservable();

  idToken$ = idToken(this.auth);
  // todo getUser (database)
  // todo getIsLoggedIn from firebase
  // use observable to get user id and if logged in 

  
  addUser(fireUid: string): Observable<{userId: number}> {
    return this.api.postUser(fireUid).pipe(map((response: Response<{userId: number}>) => {
      const data = response.data
      this.userIdSubject.next(data.userId);
       
      return data
    }))
  }


  logout(): void {
    this.auth.signOut().then(() => {
      this.router.navigate([""])
    })
  }

  constructor(private api: Api, private auth: Auth, private router: Router) { }
}
