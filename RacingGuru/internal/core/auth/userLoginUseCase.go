package auth

import (
	"RacingGuru/internal/auth"
	"RacingGuru/internal/models"
	"context"
)

type UserLoginUseCase struct {
	Repo Repo
}

func NewUserLoginUseCase(repo Repo) *UserLoginUseCase {
	return &UserLoginUseCase{Repo: repo}
}

func (uc UserLoginUseCase) Run(ctx context.Context, user models.User) (string, error) {
	user.Password = hashPassword(user.Password)
	authUser, err := uc.Repo.UserLogin(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(authUser)
	if err != nil {
		return "", err
	}
	return token, nil
}
