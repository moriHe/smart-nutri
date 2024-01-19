import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { User, UserService } from 'api/user/user.service';
import { from, map } from 'rxjs';

export const noFamilyGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService)
  const router = inject(Router)

  return from(userService.user$).pipe(
    map((user: User | null) => {
      if (user?.activeFamilyId) {
        return  router.createUrlTree(["/meine-rezepte"])
      }
      return true
    })
  )
};
