import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';

export const authGuard: CanActivateFn = (route, state) => {
  const router = inject(Router)
  const supabaseService = inject(SupabaseService)
  if (supabaseService.session?.user) {
      return true
    }
    return router.createUrlTree([""])

};
