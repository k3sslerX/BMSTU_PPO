package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetCompletionStatsUseCase struct {
	Repo Repo
	User models.User
}

func NewGetCompletionStatsUseCase(repo Repo, user models.User) *GetCompletionStatsUseCase {
	return &GetCompletionStatsUseCase{Repo: repo, User: user}
}

func (uc *GetCompletionStatsUseCase) Run(ctx context.Context) (models.SudokuCompletionStats, error) {
	return uc.Repo.GetSudokuCompletionStats(ctx, uc.User)
}
