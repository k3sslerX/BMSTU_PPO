package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetTeamMatrixUseCase struct {
	Repo Repo
	User models.User
}

func NewGetTeamMatrixUseCase(repo Repo, user models.User) *GetTeamMatrixUseCase {
	return &GetTeamMatrixUseCase{Repo: repo, User: user}
}

func (uc *GetTeamMatrixUseCase) Run(ctx context.Context) (models.MatrixTeams, error) {
	return uc.Repo.GetTeamMatrix(ctx)
}
