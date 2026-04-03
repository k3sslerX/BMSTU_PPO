package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CreateRaceUseCase struct {
	Repo Repo
	User models.User
}

func NewCreateRaceUseCase(repo Repo, user models.User) *CreateRaceUseCase {
	return &CreateRaceUseCase{Repo: repo, User: user}
}

func (uc *CreateRaceUseCase) Run(ctx context.Context, race models.Race) (models.Race, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Race{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.CreateRace(ctx, race)
}
