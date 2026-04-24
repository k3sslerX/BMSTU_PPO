.PHONY: unit-core

CORE_PATH := ./internal/core
DATA_LAYER_PATH := ./internal/repository

unit-core:
	cd RacingGuru && go test $(CORE_PATH)/... -coverprofile=$(CORE_PATH)/coverage.cov

iter-data:
	cd RacingGuru && go test $(DATA_LAYER_PATH)/... -coverprofile=$(DATA_LAYER_PATH)/coverage.cov

swag-gen:
	cd RacingGuru && make swagger

run:
	cd RacingGuru/cmd && go run .

build:
	cd RacingGuru/cmd && go build -o ../../exec/server.exe .

docker-build:
	cd RacingGuru && docker compose up --build