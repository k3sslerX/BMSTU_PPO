SWAG_BIN := $(if $(strip $(GOBIN)),$(GOBIN)/swag,$(shell go env GOPATH)/bin/swag)

SWAGGER_SRC := $(shell find cmd internal -type f -name '*.go')
SWAGGER_OUT := internal/controller/handlers/assets/swagger/swagger.yaml
SWAGGER_JSON := internal/controller/handlers/assets/swagger/swagger.json
GO_BUILD_CACHE := $(CURDIR)/.cache/go-build

.PHONY: swagger swagger-install clean-swagger docker-build-app docker-recreate-app

swagger: $(SWAGGER_OUT)

$(SWAGGER_OUT): $(SWAGGER_SRC)
	GOCACHE=$(GO_BUILD_CACHE) $(SWAG_BIN) init \
		-g swagger.go \
		-d ./internal/controller/handlers,./internal/models \
		-o ./internal/controller/handlers/assets/swagger \
		--outputTypes yaml,json

swagger-install:
	go install github.com/swaggo/swag/cmd/swag@v1.16.4

clean-swagger:
	rm -f $(SWAGGER_OUT) $(SWAGGER_JSON)

docker-build-app:
	docker compose build app

docker-recreate-app:
	docker compose up -d --no-deps --force-recreate app
