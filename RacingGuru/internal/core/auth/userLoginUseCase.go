package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type UserLoginUseCase struct {
	Repo Repo
}

func NewUserLoginUseCase(repo Repo) *UserLoginUseCase {
	return &UserLoginUseCase{Repo: repo}
}

func (uc UserLoginUseCase) Run(ctx context.Context, user models.User) (models.User, error) {
	return uc.Repo.UserLogin(ctx, user)
}
