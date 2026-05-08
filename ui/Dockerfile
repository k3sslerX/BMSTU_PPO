FROM node:22-alpine AS builder

WORKDIR /src

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:1.27-alpine AS runtime

RUN apk add --no-cache gettext jq

ENV FRONTEND_CONTAINER_PORT=80
ENV FRONTEND_SERVER_NAME=_
ENV FRONTEND_BASE_URL=
ENV API_BASE_URL=/api
ENV BACKEND_URL=http://app:8080

COPY nginx.conf.template /etc/nginx/templates/default.conf.template
COPY docker-entrypoint.d/10-write-runtime-config.sh /docker-entrypoint.d/10-write-runtime-config.sh
RUN chmod +x /docker-entrypoint.d/10-write-runtime-config.sh

COPY --from=builder /src/dist/frontend/browser /usr/share/nginx/html

EXPOSE 80
