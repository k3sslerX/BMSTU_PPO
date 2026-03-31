package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	GetDriverMatrix(context.Context) (models.MatrixDrivers, error)
	GetTeamMatrix(context.Context) (models.MatrixTeams, error)
	GetTrackMatrix(context.Context) (models.MatrixTracks, error)
}
