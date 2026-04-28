package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateUserRoleUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateUserRoleUseCase(repo Repo, user models.User) *UpdateUserRoleUseCase {
	return &UpdateUserRoleUseCase{Repo: repo, User: user}
}

func (uc *UpdateUserRoleUseCase) Run(ctx context.Context, user models.User) (models.User, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.User{}, shared.ErrorPermissionDenied
	}
	if user.Id == uc.User.Id {
		return models.User{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateUserRole(ctx, user)
}
