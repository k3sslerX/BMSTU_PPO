package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpsertRaceResultUseCase struct {
	Repo Repo
	User models.User
}

func NewUpsertRaceResultUseCase(repo Repo, user models.User) *UpsertRaceResultUseCase {
	return &UpsertRaceResultUseCase{Repo: repo, User: user}
}

func (uc *UpsertRaceResultUseCase) Run(ctx context.Context, result models.RaceResult) (models.RaceResult, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.RaceResult{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpsertRaceResult(ctx, result)
}
