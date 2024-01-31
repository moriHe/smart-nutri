import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, switchMap, take } from 'rxjs';
import { SupabaseService } from 'api/supabase.service';
import { UserService } from 'api/user/user.service';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  
  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    if (request.url.includes("secret")) {
      const secret = localStorage.getItem("secret")
      const authRequest = request.clone({
        setHeaders: {
          Secret: secret || ""
        }
      })
      return next.handle(authRequest)
    }
    return this.supabaseService.session$.pipe(
      take(1),
      switchMap(session => {
        const token = session?.access_token
        const authRequest = request.clone({
          setHeaders: {
            Authorization: token ? `Bearer ${token}` : ""
          }
        })
        return next.handle(authRequest)
      })
    )
      
    
  }
  
  constructor(private supabaseService: SupabaseService) {}
}
