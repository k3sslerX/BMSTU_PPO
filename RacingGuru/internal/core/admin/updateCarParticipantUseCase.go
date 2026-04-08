package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateCarParticipantUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateCarParticipantUseCase(repo Repo, user models.User) *UpdateCarParticipantUseCase {
	return &UpdateCarParticipantUseCase{Repo: repo, User: user}
}

func (uc *UpdateCarParticipantUseCase) Run(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.CarParticipant{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateCarParticipant(ctx, carParticipant)
}
