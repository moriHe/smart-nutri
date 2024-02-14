import { Injectable } from '@angular/core';
import { MatSnackBar, MatSnackBarRef, SimpleSnackBar } from '@angular/material/snack-bar';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SnackbarService {
  openSnackbar(message: string, action: string): Observable<void> {
  const ref: MatSnackBarRef<SimpleSnackBar> = this.snackbar.open(
      message, action, 
      {
      horizontalPosition: "start",
      verticalPosition: "bottom",
      duration: 3000
      }
    )

    return ref.onAction()
  }

  openGenericErrorSnackbar(): Observable<void> {
    const ref: MatSnackBarRef<SimpleSnackBar> = this.snackbar.open(
      "Etwas ging schief", "Ok", 
      {
      horizontalPosition: "start",
      verticalPosition: "bottom",
      duration: 3000
      }
    )

    return ref.onAction()
  }
  constructor(private snackbar: MatSnackBar) { }
}
