package admin

import (
	"RacingGuru/internal/models"
	"context"
)

type Repo interface {
	CreateDriver(context.Context, models.Driver) (models.Driver, error)
	CreateTeam(context.Context, models.Team) (models.Team, error)
	CreateTrack(context.Context, models.Track) (models.Track, error)
	CreateRace(ctx context.Context, race models.Race) (models.Race, error)
	CreateCarParticipant(context.Context, models.CarParticipant) (models.CarParticipant, error)
	UpsertRaceResult(context.Context, models.RaceResult) (models.RaceResult, error)
	UpdateCarParticipantDrivers(context.Context, models.CarParticipant) (models.CarParticipant, error)

	UpdateRace(context.Context, models.Race) (models.Race, error)
	UpdateTeam(context.Context, models.Team) (models.Team, error)
	UpdateTrack(context.Context, models.Track) (models.Track, error)
	UpdateDriver(context.Context, models.Driver) (models.Driver, error)
	UpdateCarParticipant(context.Context, models.CarParticipant) (models.CarParticipant, error)
	UpdateUserRole(context.Context, models.User) (models.User, error)
}
