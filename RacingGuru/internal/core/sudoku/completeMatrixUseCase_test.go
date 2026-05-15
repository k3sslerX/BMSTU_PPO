package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestCompleteMatrixUseCase(t *testing.T) {
	repo := &testRepo{}
	user := models.User{Role: models.RoleUser}

	err := NewCompleteMatrixUseCase(repo, user, models.SudokuMatrixTypeDrivers).Run(context.Background())
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	stats, err := NewGetCompletionStatsUseCase(repo, user).Run(context.Background())
	if err != nil {
		t.Fatalf("stats Run() error = %v", err)
	}
	if stats.DriverMatrices != 1 || stats.TeamMatrices != 0 {
		t.Fatalf("stats = %+v, want one driver completion", stats)
	}
}

func TestCompleteMatrixUseCaseRejectsUnknownMatrixType(t *testing.T) {
	err := NewCompleteMatrixUseCase(&testRepo{}, models.User{}, models.SudokuMatrixType("unknown")).Run(context.Background())
	if !errors.Is(err, shared.ErrorInvalidData) {
		t.Fatalf("Run() error = %v, want %v", err, shared.ErrorInvalidData)
	}
}
