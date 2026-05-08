package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type UserMeUseCase struct {
	repo Repo
	user models.User
}

func NewUserMeUseCase(repo Repo, user models.User) *UserMeUseCase {
	return &UserMeUseCase{repo: repo, user: user}
}

func (uc *UserMeUseCase) Run(ctx context.Context) (models.User, error) {
	return uc.repo.GetUserByID(ctx, uc.user)
}
