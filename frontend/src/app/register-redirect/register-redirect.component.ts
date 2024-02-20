import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'api/user/user.service';
import { Subject, takeUntil } from 'rxjs';

@Component({
  selector: 'app-register-redirect',
  templateUrl: './register-redirect.component.html',
  styleUrls: ['./register-redirect.component.css']
})
export class RegisterRedirectComponent {
  private destroy$ = new Subject<void>();

  ngOnInit(): void {
    this.userService.user$.pipe(
      takeUntil(this.destroy$)
    ).subscribe(user => {
      if (user) {
         this.router.navigate(["/willkommen"]);
      } else {
        this.userService.addUser().pipe(
          takeUntil(this.destroy$)
        ).subscribe(() => {
          this.userService.getUser().pipe(
            takeUntil(this.destroy$)
          ).subscribe((activeFamilyId) => {
            if (activeFamilyId) {
              this.router.navigate(["/meine-rezepte"]);
            } else {
              this.router.navigate(["/willkommen"]);
            }
          });
        });
      }
    });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  constructor(
    private router: Router,
    private userService: UserService
    ) {}
}
