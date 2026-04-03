package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateDriverUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateDriverUseCase(repo Repo, user models.User) *UpdateDriverUseCase {
	return &UpdateDriverUseCase{Repo: repo, User: user}
}

func (uc *UpdateDriverUseCase) Run(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Driver{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateDriver(ctx, driver)
}
