import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

export const adminGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);
  const notFoundRoute = router.createUrlTree(['/404']);

  return auth.sessionStatus() === 'valid' && auth.isAdmin() ? true : notFoundRoute;
};
