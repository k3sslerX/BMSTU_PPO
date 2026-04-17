package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type UserRegisterUseCase struct {
	Repo Repo
}

func NewUserRegisterUseCase(repo Repo) *UserRegisterUseCase {
	return &UserRegisterUseCase{Repo: repo}
}

func (uc UserRegisterUseCase) Run(ctx context.Context, user models.User) (models.User, error) {
	user.Password = hashPassword(user.Password)
	return uc.Repo.UserRegister(ctx, user)
}
