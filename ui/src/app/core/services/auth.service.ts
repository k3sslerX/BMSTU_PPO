import { computed, Injectable, signal } from '@angular/core';
import { AuthUserResponse, JwtClaims } from '../models/api.models';

type TokenValidationResult =
  | { status: 'valid'; claims: JwtClaims }
  | { status: 'missing' | 'expired' | 'invalid' };

export type SessionStatus = TokenValidationResult['status'];

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly storageKey = 'racingguru.auth.token';
  private readonly tokenState = signal<string | null>(this.readStoredToken());
  private readonly profileState = signal<AuthUserResponse | null>(null);

  readonly token = this.tokenState.asReadonly();
  readonly profile = this.profileState.asReadonly();
  readonly claims = computed(() => {
    const result = validateJwtToken(this.tokenState());
    return result.status === 'valid' ? result.claims : null;
  });
  readonly isLoggedIn = computed(() => this.claims() !== null);
  readonly isAdmin = computed(() => this.claims()?.role === 'admin');
  readonly role = computed(() => this.claims()?.role ?? null);
  readonly userId = computed(() => this.claims()?.user_id ?? null);
  readonly displayName = computed(() => this.profileState()?.name || 'Пользователь');
  readonly displayEmail = computed(() => this.profileState()?.email || 'Email загружается');

  constructor() {
    this.refreshTokenState();
  }

  setToken(token: string): void {
    this.tokenState.set(token);
    this.profileState.set(null);
    localStorage.setItem(this.storageKey, token);
  }

  setProfile(profile: AuthUserResponse): void {
    this.profileState.set(profile);
  }

  validToken(): string | null {
    if (this.refreshTokenState() !== 'valid') {
      return null;
    }

    return this.tokenState();
  }

  hasValidSession(): boolean {
    return this.sessionStatus() === 'valid';
  }

  sessionStatus(): SessionStatus {
    return this.refreshTokenState();
  }

  logout(): void {
    this.tokenState.set(null);
    this.profileState.set(null);
    localStorage.removeItem(this.storageKey);
  }

  private readStoredToken(): string | null {
    return localStorage.getItem(this.storageKey);
  }

  private refreshTokenState(): SessionStatus {
    const result = validateJwtToken(this.tokenState());

    if (result.status === 'expired' || result.status === 'invalid') {
      this.logout();
    }

    return result.status;
  }
}

function validateJwtToken(token: string | null): TokenValidationResult {
  if (!token) {
    return { status: 'missing' };
  }

  const segments = token.split('.');
  if (segments.length !== 3) {
    return { status: 'invalid' };
  }

  try {
    const json = decodeBase64Url(segments[1]);
    const claims = JSON.parse(json) as Partial<JwtClaims>;

    if (typeof claims.user_id !== 'string' || (claims.role !== 'admin' && claims.role !== 'user')) {
      return { status: 'invalid' };
    }

    if (typeof claims.exp !== 'number') {
      return { status: 'invalid' };
    }

    if (claims.exp <= Math.floor(Date.now() / 1000)) {
      return { status: 'expired' };
    }

    return {
      status: 'valid',
      claims: {
        user_id: claims.user_id,
        role: claims.role,
        exp: claims.exp,
        iat: claims.iat
      }
    };
  } catch {
    return { status: 'invalid' };
  }
}

function decodeBase64Url(payload: string): string {
  const normalized = payload.replace(/-/g, '+').replace(/_/g, '/');
  const padded = normalized.padEnd(Math.ceil(normalized.length / 4) * 4, '=');
  return atob(padded);
}
