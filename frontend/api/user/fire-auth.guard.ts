import { inject, runInInjectionContext } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { Auth } from "@angular/fire/auth"
import { Router } from '@angular/router';
import { UserService } from './user.service';
import { map } from 'rxjs';

export const fireAuthGuard: CanActivateFn = (route, state) => {
  const userService = inject(UserService)
  
  return userService.canActivate()
};
