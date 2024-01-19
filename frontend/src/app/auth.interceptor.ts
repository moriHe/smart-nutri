import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { SupabaseService } from 'api/supabase.service';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  
  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    const token = this.supabaseService.session?.access_token
    console.log(token)
    const authRequest = request.clone({
      setHeaders: {
        Authorization: token ? `Bearer ${token}` : ""
      }
    })
   
    return next.handle(authRequest)
      
    
  }
  
  constructor(private supabaseService: SupabaseService) {}
}
