package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CompleteMatrixUseCase struct {
	Repo       Repo
	User       models.User
	MatrixType models.SudokuMatrixType
}

func NewCompleteMatrixUseCase(repo Repo, user models.User, matrixType models.SudokuMatrixType) *CompleteMatrixUseCase {
	return &CompleteMatrixUseCase{Repo: repo, User: user, MatrixType: matrixType}
}

func (uc *CompleteMatrixUseCase) Run(ctx context.Context) error {
	if uc.MatrixType != models.SudokuMatrixTypeDrivers && uc.MatrixType != models.SudokuMatrixTypeTeams {
		return shared.ErrorInvalidData
	}

	return uc.Repo.CompleteSudokuMatrix(ctx, uc.User, uc.MatrixType)
}
