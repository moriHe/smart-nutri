import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { UserService } from 'api/user/user.service';
import { map } from 'rxjs';

export const familyGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService)
  const router = inject(Router)
  return userService.getUser().pipe(map((value => {
    if (value) {
      return true
    }
    return router.createUrlTree(["/willkommen"])
  })))
};
