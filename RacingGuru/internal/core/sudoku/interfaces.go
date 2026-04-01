package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	GetDriverMatrix(context.Context, models.MatrixDrivers) (models.MatrixDrivers, error)
	GetTeamMatrix(context.Context, models.MatrixTeams) (models.MatrixTeams, error)
}
