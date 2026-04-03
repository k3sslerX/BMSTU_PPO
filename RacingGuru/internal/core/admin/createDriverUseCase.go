package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CreateDriverUseCase struct {
	Repo Repo
	User models.User
}

func NewCreateDriverUseCase(repo Repo, user models.User) *CreateDriverUseCase {
	return &CreateDriverUseCase{Repo: repo, User: user}
}

func (uc *CreateDriverUseCase) Run(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Driver{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.CreateDriver(ctx, driver)
}
