import { Injectable } from '@angular/core';
import { Response } from 'api';
import { BehaviorSubject, Observable, catchError, map } from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment.development';

export type User = {
  id: number,
  activeFamilyId: number
}

export type UserFamily = {
  id: number
  familyId: number
  familyName: string
  role: "OWNER" | "MEMBER"
}

export type UserFamilys = UserFamily[]
@Injectable({
  providedIn: 'root'
})
export class UserService { 
  private userSubject: BehaviorSubject<User | null> = new BehaviorSubject<User | null>(null);

  user$: Observable<User | null> = this.userSubject.asObservable();

  addUser(): Observable<{userId: number}> {
    return this.http.post<Response<{userId: number}>>(`${environment.backendBaseUrl}/user`, {}).pipe(map((response: Response<{userId: number}>) => {
      return response.data
    }))
  }

  setUserNull() {
    this.userSubject.next(null)
  }


  getUser(): Observable<number> {
    return this.http.get<Response<User>>(`${environment.backendBaseUrl}/user`).pipe(map((response: Response<User>) => {
      if (response.data?.id) {
        this.userSubject.next(response.data)
      }
        return response.data?.activeFamilyId
    })
    )
  }

  getUserFamilys(): Observable<UserFamilys> {
    return this.http.get<Response<UserFamilys>>(`${environment.backendBaseUrl}/user/familys`).pipe(map((response) => response.data))
  }

  updateUserFamily(newActiveFamilyId: number): Observable<string> {
    return this.http.patch<Response<string>>(`${environment.backendBaseUrl}/user`, {newActiveFamilyId}).pipe(map((response) => {
      return response.data
    }))
  }

  getSecret(): Observable<string> {
    return this.http.get<Response<string>>(`${environment.backendBaseUrl}/secret`).pipe(map((response) => {
      return response.data
    }))
  }


  constructor(private http: HttpClient) { }
}
