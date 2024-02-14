import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, catchError, of } from 'rxjs';
import { SnackbarService } from 'services/snackbar.service';

@Injectable()
export class HttpErrorInterceptor implements HttpInterceptor {

  constructor(private snackbarServie: SnackbarService) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    return next.handle(request).pipe(
      catchError(() => {
        this.snackbarServie.openGenericErrorSnackbar()
        return []
      })
    )
  }
}
