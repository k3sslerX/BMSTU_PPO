package handlers

import (
	"RacingGuru/internal/core/admin"
	"RacingGuru/internal/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) requireUser(w http.ResponseWriter, r *http.Request) (models.User, bool) {
	user, ok := getUserFromContext(r.Context())
	if !ok {
		h.sendError(w, "unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
		return models.User{}, false
	}
	return user, true
}

func decodeJSONBody[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var payload T
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := ErrorResponse{}
		response.Error.Code = "INVALID_REQUEST"
		response.Error.Message = "invalid request"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return payload, false
	}
	return payload, true
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// CreateDriver godoc
// @Summary Create driver
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Driver true "Driver payload"
// @Success 201 {object} models.Driver
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/drivers [post]
func (h *Handler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	driver, ok := decodeJSONBody[models.Driver](w, r)
	if !ok {
		return
	}

	uc := admin.NewCreateDriverUseCase(h.Repo, user)
	created, err := uc.Run(r.Context(), driver)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateDriver godoc
// @Summary Update driver
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Driver true "Driver payload"
// @Success 200 {object} models.Driver
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/drivers [patch]
func (h *Handler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	driver, ok := decodeJSONBody[models.Driver](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateDriverUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), driver)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// UpdateUserRole godoc
// @Summary Update user role
// @Description Changes another user's role. Only users with role `admin` can call this endpoint.
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body UpdateUserRoleRequest true "User role update payload"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/users/role [patch]
func (h *Handler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	request, ok := decodeJSONBody[UpdateUserRoleRequest](w, r)
	if !ok {
		return
	}

	userID, err := uuid.Parse(request.Id)
	if err != nil {
		h.sendError(w, "invalid data", "INVALID_DATA", http.StatusBadRequest)
		return
	}

	uc := admin.NewUpdateUserRoleUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), models.User{
		Id:   userID,
		Role: request.Role,
	})
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// CreateTeam godoc
// @Summary Create team
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Team true "Team payload"
// @Success 201 {object} models.Team
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/teams [post]
func (h *Handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	team, ok := decodeJSONBody[models.Team](w, r)
	if !ok {
		return
	}

	uc := admin.NewCreateTeamUseCase(h.Repo, user)
	created, err := uc.Run(r.Context(), team)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateTeam godoc
// @Summary Update team
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Team true "Team payload"
// @Success 200 {object} models.Team
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/teams [patch]
func (h *Handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	team, ok := decodeJSONBody[models.Team](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateTeamUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), team)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// CreateTrack godoc
// @Summary Create track
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Track true "Track payload"
// @Success 201 {object} models.Track
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/tracks [post]
func (h *Handler) CreateTrack(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	track, ok := decodeJSONBody[models.Track](w, r)
	if !ok {
		return
	}

	uc := admin.NewCreateTrackUseCase(h.Repo, user)
	created, err := uc.Run(r.Context(), track)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateTrack godoc
// @Summary Update track
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Track true "Track payload"
// @Success 200 {object} models.Track
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/tracks [patch]
func (h *Handler) UpdateTrack(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	track, ok := decodeJSONBody[models.Track](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateTrackUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), track)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// CreateRace godoc
// @Summary Create race
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Race true "Race payload"
// @Success 201 {object} models.Race
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/races [post]
func (h *Handler) CreateRace(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	race, ok := decodeJSONBody[models.Race](w, r)
	if !ok {
		return
	}

	uc := admin.NewCreateRaceUseCase(h.Repo, user)
	created, err := uc.Run(r.Context(), race)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateRace godoc
// @Summary Update race
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Race true "Race payload"
// @Success 200 {object} models.Race
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/races [patch]
func (h *Handler) UpdateRace(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	race, ok := decodeJSONBody[models.Race](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateRaceUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), race)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// CreateCarParticipant godoc
// @Summary Create car participant
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.CarParticipant true "Car participant payload"
// @Success 201 {object} models.CarParticipant
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/car-participants [post]
func (h *Handler) CreateCarParticipant(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	carParticipant, ok := decodeJSONBody[models.CarParticipant](w, r)
	if !ok {
		return
	}

	uc := admin.NewCreateCarParticipantUseCase(h.Repo, user)
	created, err := uc.Run(r.Context(), carParticipant)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateCarParticipant godoc
// @Summary Update car participant
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.CarParticipant true "Car participant payload"
// @Success 200 {object} models.CarParticipant
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/car-participants [patch]
func (h *Handler) UpdateCarParticipant(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	carParticipant, ok := decodeJSONBody[models.CarParticipant](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateCarParticipantUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), carParticipant)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// UpdateCarParticipantDrivers godoc
// @Summary Update car participant drivers
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.CarParticipant true "Car participant drivers payload"
// @Success 200 {object} models.CarParticipant
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/car-participants/drivers [patch]
func (h *Handler) UpdateCarParticipantDrivers(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	carParticipant, ok := decodeJSONBody[models.CarParticipant](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpdateCarParticipantDriversUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), carParticipant)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// UpsertRaceResult godoc
// @Summary Upsert race result
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.RaceResult true "Race result payload"
// @Success 200 {object} models.RaceResult
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 403 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/race-results [post]
func (h *Handler) UpsertRaceResult(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	result, ok := decodeJSONBody[models.RaceResult](w, r)
	if !ok {
		return
	}

	uc := admin.NewUpsertRaceResultUseCase(h.Repo, user)
	updated, err := uc.Run(r.Context(), result)
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}
