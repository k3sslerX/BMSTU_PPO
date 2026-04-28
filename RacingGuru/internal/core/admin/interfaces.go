package admin

import (
	"RacingGuru/internal/models"
	"context"

	"github.com/google/uuid"
)

type Repo interface {
	ListCars(context.Context, string) ([]models.Car, error)
	ListCarParticipants(context.Context, string) ([]models.CarParticipant, error)
	ListChampionships(context.Context, string) ([]models.Championship, error)
	ListRaces(context.Context, string) ([]models.Race, error)
	ListTracks(context.Context, string) ([]models.Track, error)
	ListUsers(context.Context, string) ([]models.User, error)

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

	DeleteDriver(context.Context, uuid.UUID) error
	DeleteTeam(context.Context, uuid.UUID) error
	DeleteTrack(context.Context, uuid.UUID) error
	DeleteRace(context.Context, uuid.UUID) error
	DeleteCarParticipant(context.Context, uuid.UUID) error
}
