import { Component, inject, signal } from '@angular/core';
import { NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { catchError, finalize, of, switchMap, tap } from 'rxjs';
import { ApiService } from '../../core/services/api.service';
import { AuthService } from '../../core/services/auth.service';
import { readApiError, tokenExpiredMessage } from '../../core/utils/api-error';

interface Feedback {
  kind: 'success' | 'error';
  title: string;
  text: string;
}

@Component({
  selector: 'app-auth-page',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './auth-page.component.html'
})
export class AuthPageComponent {
  private readonly fb = inject(NonNullableFormBuilder);
  private readonly api = inject(ApiService);
  private readonly auth = inject(AuthService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  protected readonly loginBusy = signal(false);
  protected readonly registerBusy = signal(false);
  protected readonly loginFeedback = signal<Feedback | null>(null);
  protected readonly registerFeedback = signal<Feedback | null>(null);

  protected readonly loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required]]
  });

  protected readonly registerForm = this.fb.group({
    name: [''],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(3)]]
  });

  constructor() {
    if (this.route.snapshot.queryParamMap.get('session') === 'expired') {
      this.loginFeedback.set({
        kind: 'error',
        title: 'Сессия истекла',
        text: tokenExpiredMessage
      });
    }
  }

  protected submitLogin(): void {
    if (this.loginForm.invalid) {
      this.loginForm.markAllAsTouched();
      this.loginFeedback.set({
        kind: 'error',
        title: 'Проверьте поля входа',
        text: 'Email и пароль обязательны. Email должен быть в формате user@example.com.'
      });
      return;
    }

    this.loginFeedback.set(null);
    this.loginBusy.set(true);

    const { email, password } = this.loginForm.getRawValue();

    this.api
      .login({ email, password })
      .pipe(
        tap(({ token }) => {
          this.auth.setToken(token);
        }),
        switchMap(() =>
          this.api.getMe().pipe(
            catchError(() => {
              return of(null);
            })
          )
        ),
        finalize(() => this.loginBusy.set(false))
      )
      .subscribe({
        next: (profile) => {
          if (profile) {
            this.auth.setProfile(profile);
          }
          this.loginFeedback.set({
            kind: 'success',
            title: 'Сессия открыта',
            text: 'JWT сохранён локально. Переходим в account, где можно менять пароль и управлять избранным.'
          });
          void this.router.navigateByUrl('/account');
        },
        error: (error) => {
          this.loginFeedback.set({
            kind: 'error',
            title: 'Логин не удался',
            text: readApiError(error)
          });
        }
      });
  }

  protected submitRegister(): void {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched();
      this.registerFeedback.set({
        kind: 'error',
        title: 'Проверьте регистрацию',
        text: 'Email и пароль обязательны, пароль должен быть не короче 3 символов.'
      });
      return;
    }

    this.registerFeedback.set(null);
    this.registerBusy.set(true);

    const { name, email, password } = this.registerForm.getRawValue();

    this.api
      .register({
        name: name.trim(),
        email,
        password
      })
      .pipe(finalize(() => this.registerBusy.set(false)))
      .subscribe({
        next: (response) => {
          this.registerFeedback.set({
            kind: 'success',
            title: 'Пользователь создан',
            text: `${response.email} зарегистрирован с ролью ${response.role}. Теперь можно выполнить login.`
          });
          this.registerForm.patchValue({ password: '' });
        },
        error: (error) => {
          this.registerFeedback.set({
            kind: 'error',
            title: 'Регистрация не удалась',
            text: readApiError(error)
          });
        }
      });
  }
}
