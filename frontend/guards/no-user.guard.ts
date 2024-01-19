import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { SupabaseService } from 'api/supabase.service';
import { from, map } from 'rxjs';

export const noUserGuard: CanActivateFn = (route, state) => {
  const router = inject(Router)
  const supabaseService = inject(SupabaseService)

  return from(supabaseService.session$).pipe(
    map((session) => {
      if (session?.user) {
        return router.createUrlTree(["/meine-rezepte"])
      }
      return true
    })
  )
};
