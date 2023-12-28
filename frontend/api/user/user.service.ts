import { Injectable } from '@angular/core';
import { Api, Response } from 'api';
import { BehaviorSubject, Observable, Subscription, map } from 'rxjs';
import { Auth, idToken, authState, User } from '@angular/fire/auth';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private userIdSubject = new BehaviorSubject<number | null>(null);
  userId$ = this.userIdSubject.asObservable();
  userIdSubscription = this.userId$.subscribe((userId: number | null) => userId)

  idToken$ = idToken(this.auth);
  idTokenSubscription = this.idToken$.subscribe((token: string | null) => token);
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

  ngOnDestroy() {
    this.idTokenSubscription.unsubscribe();
    this.userIdSubscription.unsubscribe();
  }

  constructor(private api: Api, private auth: Auth) { }
}
