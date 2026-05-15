package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	GetDriverMatrix(context.Context, models.MatrixDrivers) (models.MatrixDrivers, error)
	GetTeamMatrix(context.Context, models.MatrixTeams) (models.MatrixTeams, error)
	CompleteSudokuMatrix(context.Context, models.User, models.SudokuMatrixType) error
	GetSudokuCompletionStats(context.Context, models.User) (models.SudokuCompletionStats, error)
}
