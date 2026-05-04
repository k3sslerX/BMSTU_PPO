import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

export const authGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);
  const sessionStatus = auth.sessionStatus();

  if (sessionStatus === 'valid') {
    return true;
  }

  if (sessionStatus === 'expired') {
    return router.createUrlTree(['/auth'], { queryParams: { session: 'expired' } });
  }

  return router.createUrlTree(['/auth']);
};
