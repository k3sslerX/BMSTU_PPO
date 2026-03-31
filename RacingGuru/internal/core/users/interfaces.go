package users

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	ToggleFavouriteDriver(context.Context, models.User, models.Driver) error
	ToggleFavouriteTeam(context.Context, models.User, models.Team) error

	GetDriverByName(context.Context, string) (models.Driver, error)
	GetTeamByName(context.Context, string) (models.Team, error)
}
