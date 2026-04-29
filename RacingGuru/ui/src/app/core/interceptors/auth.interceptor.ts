import { HttpErrorResponse, HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { throwError } from 'rxjs';
import { AuthService } from '../services/auth.service';

export const authInterceptor: HttpInterceptorFn = (request, next) => {
  const auth = inject(AuthService);
  const router = inject(Router);
  const sessionStatus = auth.sessionStatus();

  if (sessionStatus === 'expired' && requiresAuth(request.url)) {
    void router.navigate(['/auth'], { queryParams: { session: 'expired' } });

    return throwError(
      () =>
        new HttpErrorResponse({
          error: { error: { code: 'TOKEN_EXPIRED', message: 'token expired' } },
          status: 401,
          statusText: 'Unauthorized',
          url: request.url
        })
    );
  }

  const token = sessionStatus === 'valid' ? auth.validToken() : null;

  if (!token) {
    return next(request);
  }

  return next(
    request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    })
  );
};

function requiresAuth(url: string): boolean {
  const path = url.replace(/^https?:\/\/[^/]+/, '');

  return (
    path.startsWith('/api/admin') ||
    path.startsWith('/api/sudoku') ||
    path.startsWith('/api/users') ||
    path === '/api/change-password'
  );
}
