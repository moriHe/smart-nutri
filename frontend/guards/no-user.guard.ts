import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';

export const noUserGuard: CanActivateFn = (route, state) => {
  const router = inject(Router)
  const supabaseService = inject(SupabaseService)

  if (supabaseService.session?.user) {
    return router.createUrlTree(["/meine-rezepte"])
  }
  return true
};
