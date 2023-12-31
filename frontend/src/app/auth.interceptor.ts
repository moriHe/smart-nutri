import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, filter, switchMap } from 'rxjs';
import {Auth, idToken} from "@angular/fire/auth"

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  
  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    return idToken(this.auth).pipe(
      filter((token: string | null): token is string => token !== null),
      switchMap((token: string) => {
        const authRequest = request.clone({
          setHeaders: {
            Authorization: token ? `Bearer ${token}` : ""
          }
        })
        return next.handle(authRequest)
      })
    )
  }
  
  constructor(private auth: Auth) {}
}
