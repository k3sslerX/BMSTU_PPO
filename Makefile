.PHONY: unit-core

CORE_PATH := ./internal/core
DATA_LAYER_PATH := ./internal/repository

unit-core:
	cd RacingGuru && go test $(CORE_PATH)/... -coverprofile=$(CORE_PATH)/coverage.cov

iter-data:
	cd RacingGuru && go test $(DATA_LAYER_PATH)/... -coverprofile=$(DATA_LAYER_PATH)/coverage.cov

swag-gen:
	cd RacingGuru && make swagger