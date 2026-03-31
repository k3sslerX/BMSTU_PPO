package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetTrackMatrixUseCase struct {
	Repo Repo
	User models.User
}

func NewGetTrackMatrixUseCase(repo Repo, user models.User) *GetTrackMatrixUseCase {
	return &GetTrackMatrixUseCase{Repo: repo, User: user}
}

func (uc *GetTrackMatrixUseCase) Run(ctx context.Context) (models.MatrixTracks, error) {
	return uc.Repo.GetTrackMatrix(ctx)
}
