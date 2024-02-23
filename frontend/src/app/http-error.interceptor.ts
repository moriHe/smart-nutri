import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { SnackbarService } from 'services/snackbar.service';
import { Router } from '@angular/router';

@Injectable()
export class HttpErrorInterceptor implements HttpInterceptor {

  constructor(private snackbarService: SnackbarService) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    if (this.shouldShowGenericError(request.method, request.url)) {
      return next.handle(request).pipe(
        catchError(() => {
          // Handle errors for specific endpoints
          this.snackbarService.openGenericErrorSnackbar()
          return []
        })
      );
    } else {
      // For other requests, simply return the next handle
      return next.handle(request);
    }
  }

  shouldShowGenericError(method: string, url: string): boolean {
    if (url.includes("user") && (method === "GET" || method === "DELETE")) {
      return false
    }
    return true
  }
}
