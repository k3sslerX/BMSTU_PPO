package stats

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	GetDriverStats(context.Context, models.Driver) (models.DriverStats, error)
	GetTeamStats(context.Context, models.Team) (models.TeamStats, error)

	GetDriverByName(context.Context, string) (models.Driver, error)
	GetTeamByName(context.Context, string) (models.Team, error)
}
