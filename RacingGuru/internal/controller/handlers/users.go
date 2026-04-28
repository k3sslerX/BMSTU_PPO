package handlers

import (
	"RacingGuru/internal/core/users"
	"RacingGuru/internal/models"
	"encoding/json"
	"net/http"
)

// ListFavourites godoc
// @Summary List favourite drivers and teams
// @Tags users
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} FavouritesResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/favourites [get]
func (h *Handler) ListFavourites(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := users.NewListFavouritesUseCase(h.Repo, user)
	drivers, teams, err := uc.Run(r.Context())
	if err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(FavouritesResponse{
		FavouriteDrivers: drivers,
		FavouriteTeams:   teams,
	})
}

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
	driver, ok := decodeJSONBody[models.Driver](h, w, r)
	if !ok {
		return
	}

	uc := users.NewToggleFavouriteDriverUseCase(h.Repo, user)
	if err := uc.Run(r.Context(), driver); err != nil {
		h.sendErrorExpanded(w, r, err)
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
	team, ok := decodeJSONBody[models.Team](h, w, r)
	if !ok {
		return
	}

	uc := users.NewToggleFavouriteTeamUseCase(h.Repo, user)
	if err := uc.Run(r.Context(), team); err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
