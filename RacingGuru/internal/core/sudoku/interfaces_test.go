package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
	returnError bool
	completions map[models.SudokuMatrixType]int64
}

func (repo *testRepo) GetDriverMatrix(ctx context.Context, matrix models.MatrixDrivers) (models.MatrixDrivers, error) {
	if repo.returnError {
		return models.MatrixDrivers{}, shared.ErrorNotFound
	}
	return matrix, nil
}

func (repo *testRepo) GetTeamMatrix(ctx context.Context, matrix models.MatrixTeams) (models.MatrixTeams, error) {
	if repo.returnError {
		return models.MatrixTeams{}, shared.ErrorNotFound
	}
	return matrix, nil
}

func (repo *testRepo) CompleteSudokuMatrix(ctx context.Context, user models.User, matrixType models.SudokuMatrixType) error {
	if repo.returnError {
		return shared.ErrorNotFound
	}
	if repo.completions == nil {
		repo.completions = make(map[models.SudokuMatrixType]int64)
	}
	repo.completions[matrixType] = 1
	return nil
}

func (repo *testRepo) GetSudokuCompletionStats(ctx context.Context, user models.User) (models.SudokuCompletionStats, error) {
	if repo.returnError {
		return models.SudokuCompletionStats{}, shared.ErrorNotFound
	}
	return models.SudokuCompletionStats{
		DriverMatrices: repo.completions[models.SudokuMatrixTypeDrivers],
		TeamMatrices:   repo.completions[models.SudokuMatrixTypeTeams],
	}, nil
}
