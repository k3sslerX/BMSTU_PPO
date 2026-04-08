package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateCarParticipantDriversUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateCarParticipantDriversUseCase(repo Repo, user models.User) *UpdateCarParticipantDriversUseCase {
	return &UpdateCarParticipantDriversUseCase{Repo: repo, User: user}
}

func (uc *UpdateCarParticipantDriversUseCase) Run(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.CarParticipant{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateCarParticipantDrivers(ctx, carParticipant)
}
