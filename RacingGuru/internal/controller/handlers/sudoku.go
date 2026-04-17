package handlers

import (
	"RacingGuru/internal/core/sudoku"
	"net/http"
)

// SudokuDrivers godoc
// @Summary Get driver sudoku matrix
// @Tags sudoku
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.MatrixDrivers
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sudoku/drivers [get]
func (h *Handler) SudokuDrivers(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := sudoku.NewGetDriverMatrixUseCase(h.Repo, user)
	matrix, err := uc.Run(r.Context())
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, matrix)
}

// SudokuTeams godoc
// @Summary Get team sudoku matrix
// @Tags sudoku
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.MatrixTeams
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sudoku/teams [get]
func (h *Handler) SudokuTeams(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := sudoku.NewGetTeamMatrixUseCase(h.Repo, user)
	matrix, err := uc.Run(r.Context())
	if err != nil {
		h.sendErrorExpanded(w, err)
		return
	}

	writeJSON(w, http.StatusOK, matrix)
}
