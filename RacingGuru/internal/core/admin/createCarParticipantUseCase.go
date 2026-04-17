package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CreateCarParticipantUseCase struct {
	Repo Repo
	User models.User
}

func NewCreateCarParticipantUseCase(repo Repo, user models.User) *CreateCarParticipantUseCase {
	return &CreateCarParticipantUseCase{Repo: repo, User: user}
}

func (uc *CreateCarParticipantUseCase) Run(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.CarParticipant{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.CreateCarParticipant(ctx, carParticipant)
}
