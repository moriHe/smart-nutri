import { Injectable } from '@angular/core';
import { Response } from 'api';
import { BehaviorSubject, Observable, map } from 'rxjs';
import { Auth } from '@angular/fire/auth';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

type User ={
  id: number,
  activeFamilyId: number
}
@Injectable({
  providedIn: 'root'
})
export class UserService { 
  addUser(): Observable<{userId: number}> {
    return this.http.post<Response<{userId: number}>>("http://localhost:8080/user", {}).pipe(map((response: Response<{userId: number}>) => {
      return response.data
    }))
  }

  getUser(): Observable<number> {
    return this.http.get<Response<User>>("http://localhost:8080/user").pipe(map((response: Response<User>) => {
        return response.data.activeFamilyId
    }))
  }

  logout(): void {
    this.auth.signOut().then(() => {
      this.router.navigate([""])
    })
  }

  constructor(private http: HttpClient, private auth: Auth, private router: Router) { }
}
