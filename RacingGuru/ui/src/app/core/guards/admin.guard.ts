import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

export const adminGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);
  const sessionStatus = auth.sessionStatus();

  if (sessionStatus === 'expired') {
    return router.createUrlTree(['/auth'], { queryParams: { session: 'expired' } });
  }

  if (sessionStatus !== 'valid') {
    return router.createUrlTree(['/auth']);
  }

  return auth.isAdmin() ? true : router.createUrlTree(['/account']);
};
