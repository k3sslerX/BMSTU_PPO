.PHONY: unit-core

CORE_PATH := ./internal/core

unit-core:
	cd RacingGuru && go test $(CORE_PATH)/... -coverprofile=$(CORE_PATH)/coverage.cov