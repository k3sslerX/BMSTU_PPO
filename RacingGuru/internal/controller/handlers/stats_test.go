package handlers

import (
	"RacingGuru/internal/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatsDriverRequiresDriverID(t *testing.T) {
	handler := newTestHandler(t)
	request := httptest.NewRequest(http.MethodGet, "/stats/driver", nil)
	response := httptest.NewRecorder()

	handler.StatsDriver(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("StatsDriver() status = %d, want %d", response.Code, http.StatusBadRequest)
	}
	if !strings.Contains(response.Body.String(), "INVALID_REQUEST") {
		t.Fatalf("StatsDriver() body = %q, want INVALID_REQUEST", response.Body.String())
	}
}

func TestStatsTeamRequiresTeamID(t *testing.T) {
	handler := newTestHandler(t)
	request := httptest.NewRequest(http.MethodGet, "/stats/team", nil)
	response := httptest.NewRecorder()

	handler.StatsTeam(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("StatsTeam() status = %d, want %d", response.Code, http.StatusBadRequest)
	}
	if !strings.Contains(response.Body.String(), "INVALID_REQUEST") {
		t.Fatalf("StatsTeam() body = %q, want INVALID_REQUEST", response.Body.String())
	}
}

func newTestHandler(t *testing.T) *Handler {
	t.Helper()

	appLogger, err := logger.New("debug", io.Discard)
	if err != nil {
		t.Fatalf("logger.New() error = %v", err)
	}

	return NewHandler(nil, appLogger)
}
