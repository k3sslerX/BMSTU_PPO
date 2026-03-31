package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetDriverMatrixUseCase struct {
	Repo Repo
	User models.User
}

func NewGetDriverMatrixUseCase(repo Repo, user models.User) *GetDriverMatrixUseCase {
	return &GetDriverMatrixUseCase{Repo: repo, User: user}
}

func (uc *GetDriverMatrixUseCase) Run(ctx context.Context) (models.MatrixDrivers, error) {
	return uc.Repo.GetDriverMatrix(ctx)
}
