package auth

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	UserLogin(context.Context, models.User) (models.User, error)
	UserRegister(context.Context, models.User) (models.User, error)
	UserChangePassword(context.Context, models.User, string) error
}
