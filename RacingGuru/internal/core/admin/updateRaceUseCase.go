package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateRaceUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateRaceUseCase(repo Repo, user models.User) *UpdateRaceUseCase {
	return &UpdateRaceUseCase{Repo: repo, User: user}
}

func (uc *UpdateRaceUseCase) Run(ctx context.Context, race models.Race) (models.Race, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Race{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateRace(ctx, race)
}
