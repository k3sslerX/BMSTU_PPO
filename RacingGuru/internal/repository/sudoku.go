package repository

import (
	"RacingGuru/internal/models"
	"context"
)

func (r *Repository) GetDriverMatrix(ctx context.Context, matrix models.MatrixDrivers) (models.MatrixDrivers, error) {
	return matrix, nil
}

func (r *Repository) GetTeamMatrix(ctx context.Context, matrix models.MatrixTeams) (models.MatrixTeams, error) {
	return matrix, nil
}
