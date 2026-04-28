package stats

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	GetDriverStats(context.Context, models.Driver) (models.DriverStats, error)
	GetTeamStats(context.Context, models.Team) (models.TeamStats, error)
	ListDrivers(context.Context, string) ([]models.Driver, error)
	ListTeams(context.Context, string) ([]models.Team, error)

	GetDriverByName(context.Context, string) (models.Driver, error)
	GetTeamByName(context.Context, string) (models.Team, error)
}
