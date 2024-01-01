import { Injectable, signal } from '@angular/core';
import { Response } from 'api';
import { BehaviorSubject, Observable, finalize, map, of, switchMap, take } from 'rxjs';
import { Auth, idToken, authState, User } from '@angular/fire/auth';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

export type DbUser = {
  id: number,
  activeFamilyId: number | null
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private isInitializedSubject = new BehaviorSubject<boolean>(false);
  isInitialized$ = this.isInitializedSubject.asObservable();
   
  user = authState(this.auth).pipe(
    take(1),
    switchMap((authUser) => {
    if (authUser) {
      return this.http.get<Response<DbUser>>("http://localhost:8080/user")
    }
    return of(null)
    }),
    finalize(() => {
      this.isInitializedSubject.next(true)
    }))

    canActivate() {
      return this.user.pipe((map(user => {
        if (user) {
          return true
        }
        return this.router.createUrlTree([""])
      })))
    }


  private userIdSubject = new BehaviorSubject<number | null>(null);
  userId$ = this.userIdSubject.asObservable();
  
  addUser(fireUid: string): Observable<{userId: number}> {
    return this.http.post<Response<{userId: number}>>("http://localhost:8080/user", {
      fireUid
  }).pipe(map((response: Response<{userId: number}>) => {
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

  constructor(private http: HttpClient, private auth: Auth, private router: Router) { }
}
