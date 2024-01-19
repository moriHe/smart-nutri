import { Injectable } from '@angular/core';
import { Response } from 'api';
import { BehaviorSubject, Observable, map } from 'rxjs';
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
    // TODO should store user once to not trigger on each authguard
    return this.http.get<Response<User>>("http://localhost:8080/user").pipe(map((response: Response<User>) => {
        return response.data.activeFamilyId
    }))
  }
  

  constructor(private http: HttpClient, private router: Router) { }
}
