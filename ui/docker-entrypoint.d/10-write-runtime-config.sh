#!/bin/sh
set -eu

CONFIG_PATH=/usr/share/nginx/html/config.js

jq -n \
  --arg apiBaseUrl "${API_BASE_URL:-/api}" \
  --arg frontendBaseUrl "${FRONTEND_BASE_URL:-}" \
  '{apiBaseUrl: $apiBaseUrl, frontendBaseUrl: $frontendBaseUrl}' \
  | sed '1s/^/window.__RACING_GURU_CONFIG__ = /;$s/$/;/' > "$CONFIG_PATH"

