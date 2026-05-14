package handlers

import (
	"RacingGuru/internal/core/sudoku"
	"RacingGuru/internal/models"
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
		h.sendErrorExpanded(w, r, err)
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
		h.sendErrorExpanded(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, matrix)
}

// CompleteSudokuDrivers godoc
// @Summary Mark driver sudoku matrix as completed
// @Tags sudoku
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 204
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sudoku/drivers/completions [post]
func (h *Handler) CompleteSudokuDrivers(w http.ResponseWriter, r *http.Request) {
	h.completeSudokuMatrix(w, r, models.SudokuMatrixTypeDrivers)
}

// CompleteSudokuTeams godoc
// @Summary Mark team sudoku matrix as completed
// @Tags sudoku
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 204
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sudoku/teams/completions [post]
func (h *Handler) CompleteSudokuTeams(w http.ResponseWriter, r *http.Request) {
	h.completeSudokuMatrix(w, r, models.SudokuMatrixTypeTeams)
}

func (h *Handler) completeSudokuMatrix(w http.ResponseWriter, r *http.Request, matrixType models.SudokuMatrixType) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := sudoku.NewCompleteMatrixUseCase(h.Repo, user, matrixType)
	if err := uc.Run(r.Context()); err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SudokuCompletions godoc
// @Summary Get sudoku completion counts
// @Tags sudoku
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.SudokuCompletionStats
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sudoku/completions [get]
func (h *Handler) SudokuCompletions(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}

	uc := sudoku.NewGetCompletionStatsUseCase(h.Repo, user)
	stats, err := uc.Run(r.Context())
	if err != nil {
		h.sendErrorExpanded(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, stats)
}
