package handlers

import (
	"RacingGuru/internal/core/users"
	"RacingGuru/internal/models"
	"net/http"
)

// ToggleFavouriteDriver godoc
// @Summary Toggle favourite driver
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Driver true "Driver payload"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/favourite-driver [post]
func (h *Handler) ToggleFavouriteDriver(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	driver, ok := decodeJSONBody[models.Driver](w, r)
	if !ok {
		return
	}

	uc := users.NewToggleFavouriteDriverUseCase(h.Repo, user)
	if err := uc.Run(r.Context(), driver); err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ToggleFavouriteTeam godoc
// @Summary Toggle favourite team
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body models.Team true "Team payload"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/favourite-team [post]
func (h *Handler) ToggleFavouriteTeam(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	team, ok := decodeJSONBody[models.Team](w, r)
	if !ok {
		return
	}

	uc := users.NewToggleFavouriteTeamUseCase(h.Repo, user)
	if err := uc.Run(r.Context(), team); err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
