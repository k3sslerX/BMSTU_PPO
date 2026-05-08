package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	HasAdmin(context.Context) (bool, error)
	GetUserByID(context.Context, models.User) (models.User, error)
	UserLogin(context.Context, models.User) (models.User, error)
	UserRegister(context.Context, models.User) (models.User, error)
	UserChangePassword(context.Context, models.User, string) error
}
