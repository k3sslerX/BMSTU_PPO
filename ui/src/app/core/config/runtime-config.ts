export interface RuntimeConfig {
  apiBaseUrl: string;
  frontendBaseUrl: string;
}

declare global {
  interface Window {
    __RACING_GURU_CONFIG__?: Partial<RuntimeConfig>;
  }
}

export function getRuntimeConfig(): RuntimeConfig {
  const config = window.__RACING_GURU_CONFIG__ ?? {};

  return {
    apiBaseUrl: normalizeBaseUrl(config.apiBaseUrl, '/api'),
    frontendBaseUrl: normalizeBaseUrl(config.frontendBaseUrl, '')
  };
}

function normalizeBaseUrl(value: string | undefined, fallback: string): string {
  const normalized = value?.trim().replace(/\/+$/, '');

  return normalized || fallback;
}
