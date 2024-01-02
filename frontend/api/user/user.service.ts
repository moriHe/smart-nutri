import { Injectable, signal } from '@angular/core';
import { Response } from 'api';
import { BehaviorSubject, Observable, finalize, map, take } from 'rxjs';
import { Auth, authState } from '@angular/fire/auth';
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
  addUser(fireUid: string): Observable<{userId: number}> {
    return this.http.post<Response<{userId: number}>>("http://localhost:8080/user", {
      fireUid
  }).pipe(map((response: Response<{userId: number}>) => {
      return response.data
    }))
  }


  logout(): void {
    this.auth.signOut().then(() => {
      this.router.navigate([""])
    })
  }

  constructor(private http: HttpClient, private auth: Auth, private router: Router) { }
}
