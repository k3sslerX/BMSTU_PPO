package bootstrap

import (
	"RacingGuru/internal/core/auth"
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
)

const (
	defaultAdminName     = "admin"
	defaultAdminLogin    = "admin"
	defaultAdminPassword = "admin"
)

type authRepo interface {
	HasAdmin(context.Context) (bool, error)
	UserLogin(context.Context, models.User) (models.User, error)
	UserRegister(context.Context, models.User) (models.User, error)
	UserChangePassword(context.Context, models.User, string) error
}

func EnsureAdmin(ctx context.Context, repo authRepo) error {
	hasAdmin, err := repo.HasAdmin(ctx)
	if err != nil {
		return err
	}
	if hasAdmin {
		return nil
	}

	_, err = auth.NewUserRegisterUseCase(repo).Run(ctx, models.User{
		Name:     defaultAdminName,
		Email:    defaultAdminLogin,
		Password: defaultAdminPassword,
		Role:     models.RoleAdmin,
	})
	if err != nil && !errors.Is(err, shared.ErrorUserAlreadyExists) {
		return err
	}

	return nil
}
