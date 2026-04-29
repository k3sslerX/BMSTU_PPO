import { HttpErrorResponse } from '@angular/common/http';
import { ApiErrorEnvelope } from '../models/api.models';

export const tokenExpiredMessage = 'Сессия истекла. Войдите заново, чтобы продолжить.';

export function readApiError(error: unknown): string {
  if (error instanceof HttpErrorResponse) {
    const envelope = error.error as ApiErrorEnvelope | null;
    const code = envelope?.error?.code;
    const message = envelope?.error?.message?.trim();
    const rawError = typeof error.error === 'string' ? error.error : '';

    if (code === 'TOKEN_EXPIRED') {
      return tokenExpiredMessage;
    }

    if (message) {
      return message;
    }

    if (error.status === 0) {
      return 'API недоступно. Проверьте, что Go-сервер запущен на 8080 порту.';
    }

    if (error.status >= 500 && (rawError.includes('proxy') || rawError.includes('ECONNREFUSED'))) {
      return 'Backend не отвечает. Запустите Go-сервер на http://localhost:8080 и повторите вход.';
    }

    return `Запрос завершился с ошибкой ${error.status}.`;
  }

  if (error instanceof Error && error.message.trim()) {
    return error.message;
  }

  return 'Не удалось выполнить запрос.';
}
