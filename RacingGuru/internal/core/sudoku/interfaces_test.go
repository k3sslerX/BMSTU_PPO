package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
	returnError bool
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
