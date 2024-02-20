import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, catchError, of } from 'rxjs';
import { SnackbarService } from 'services/snackbar.service';
import { Router } from '@angular/router';

@Injectable()
export class HttpErrorInterceptor implements HttpInterceptor {

  constructor(private snackbarServie: SnackbarService, private router: Router) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    return next.handle(request).pipe(
      catchError(() => {
       if (!this.router.url.includes("registrieren/redirect")) {
         this.snackbarServie.openGenericErrorSnackbar()
       }
        return []
      })
    )
  }
}
