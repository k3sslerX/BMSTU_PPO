package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
	returnError bool
}

func (repo *testRepo) GetDriverMatrix(ctx context.Context) (models.MatrixDrivers, error) {
	if repo.returnError {
		return models.MatrixDrivers{}, shared.ErrorNotFound
	}
	return models.MatrixDrivers{}, nil
}

func (repo *testRepo) GetTeamMatrix(ctx context.Context) (models.MatrixTeams, error) {
	if repo.returnError {
		return models.MatrixTeams{}, shared.ErrorNotFound
	}
	return models.MatrixTeams{}, nil
}

func (repo *testRepo) GetTrackMatrix(ctx context.Context) (models.MatrixTracks, error) {
	if repo.returnError {
		return models.MatrixTracks{}, shared.ErrorNotFound
	}
	return models.MatrixTracks{}, nil
}
