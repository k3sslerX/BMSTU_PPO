package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type UserChangePasswordUseCase struct {
	repo Repo
	user models.User
}

func NewUserChangePasswordUseCase(repo Repo, user models.User) *UserChangePasswordUseCase {
	return &UserChangePasswordUseCase{repo: repo, user: user}
}

func (uc *UserChangePasswordUseCase) Run(ctx context.Context, pwd string) error {
	return uc.repo.UserChangePassword(ctx, uc.user, hashPassword(pwd))
}
