package handlers

import (
	"RacingGuru/internal/core/stats"
	"RacingGuru/internal/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// StatsDriver godoc
// @Summary Driver statistics
// @Tags stats
// @Produce json
// @Param driver_id query string true "Driver UUID" format(uuid)
// @Success 200 {object} DriverStatsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stats/driver [get]
func (h *Handler) StatsDriver(w http.ResponseWriter, r *http.Request) {
	//user, ok := getUserFromContext(r.Context())
	//if !ok {
	//	h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusUnauthorized)
	//	return
	//}
	var driver models.Driver
	driverID := r.URL.Query().Get("driver_id")
	if driverID != "" {
		parsedID, err := uuid.Parse(driverID)
		if err != nil {
			h.sendError(w, "invalid driver_id", "INVALID_REQUEST", http.StatusBadRequest)
			return
		}
		driver.Id = parsedID
	} else {
		return
	}
	uc := stats.NewDriverStatsUseCase(h.Repo)
	dStats, err := uc.Run(r.Context(), driver)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	wrapper := DriverStatsResponse{Stats: dStats}
	_ = json.NewEncoder(w).Encode(wrapper)
}

// StatsTeam godoc
// @Summary Team statistics
// @Tags stats
// @Produce json
// @Param team_id query string true "Team UUID" format(uuid)
// @Success 200 {object} TeamStatsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stats/team [get]
func (h *Handler) StatsTeam(w http.ResponseWriter, r *http.Request) {
	//user, ok := getUserFromContext(r.Context())
	//if !ok {
	//	h.sendError(w, "invalid request", "INVALID_REQUEST", http.StatusUnauthorized)
	//	return
	//}
	var team models.Team
	teamID := r.URL.Query().Get("team_id")
	if teamID != "" {
		parsedID, err := uuid.Parse(teamID)
		if err != nil {
			h.sendError(w, "invalid team_id", "INVALID_REQUEST", http.StatusBadRequest)
			return
		}
		team.Id = parsedID
	} else {
		return
	}
	uc := stats.NewTeamStatsUseCase(h.Repo)
	tStats, err := uc.Run(r.Context(), team)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	wrapper := TeamStatsResponse{Stats: tStats}
	_ = json.NewEncoder(w).Encode(wrapper)
}
