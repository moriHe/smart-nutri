import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import {Auth, authState} from '@angular/fire/auth'
import { map, switchMap } from 'rxjs';
import { SupabaseService } from 'api/supabase.service';

export const authGuard: CanActivateFn = (route, state) => {
  const auth = inject(Auth)
  const router = inject(Router)
  const supabaseService = inject(SupabaseService)
  if (supabaseService.session?.user) {
      return true
    }
    return router.createUrlTree([""])

};
