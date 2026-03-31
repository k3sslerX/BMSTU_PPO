package admin

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	CreateDriver(context.Context, models.Driver) (models.Driver, error)
	CreateTeam(context.Context, models.Team) (models.Team, error)
	CreateRace(ctx context.Context, race models.Race) (models.Race, error)

	UpdateRace(context.Context, models.Race) (models.Race, error)
	UpdateTeam(context.Context, models.Team) (models.Team, error)
	UpdateDriver(context.Context, models.Driver) (models.Driver, error)
}
