FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/racingguru ./cmd

FROM alpine:3.21 AS runtime

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /out/racingguru /app/racingguru

ENV SERVER_ADDR=:8080
ENV DBMS=postgres
ENV APP_POSTGRES_DATABASE_URL=
ENV APP_MONGO_DATABASE_URL=
ENV APP_MONGO_DATABASE_NAME=
ENV LOG_LEVEL=
ENV LOG_FILE=

USER app

EXPOSE 8080

CMD ["/app/racingguru"]
