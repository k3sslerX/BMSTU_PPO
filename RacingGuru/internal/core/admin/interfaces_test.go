package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type testRepo struct {
}

func (repo *testRepo) ListCars(ctx context.Context, query string) ([]models.Car, error) {
	return []models.Car{{Id: uuid.MustParse("88888888-8888-8888-8888-888888888888"), Model: "testCar", Year: 2024}}, nil
}

func (repo *testRepo) ListCarParticipants(ctx context.Context, query string) ([]models.CarParticipant, error) {
	return []models.CarParticipant{{Id: uuid.MustParse("66666666-6666-6666-6666-666666666666"), Number: "7"}}, nil
}

func (repo *testRepo) ListChampionships(ctx context.Context, query string) ([]models.Championship, error) {
	return []models.Championship{{Id: uuid.MustParse("99999999-9999-9999-9999-999999999999"), Year: 2024}}, nil
}

func (repo *testRepo) ListRaces(ctx context.Context, query string) ([]models.Race, error) {
	return []models.Race{{Id: uuid.MustParse("33333333-3333-3333-3333-333333333333"), Name: "testRace"}}, nil
}

func (repo *testRepo) ListTracks(ctx context.Context, query string) ([]models.Track, error) {
	return []models.Track{{Id: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "testTrack"}}, nil
}

func (repo *testRepo) ListUsers(ctx context.Context, query string) ([]models.User, error) {
	return []models.User{{Id: uuid.MustParse("77777777-7777-7777-7777-777777777777"), Name: "testUser", Email: "test@example.com", Role: models.RoleUser}}, nil
}

func (repo *testRepo) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Name == "testName" {
		return models.Driver{Id: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Name == "testName" {
		return models.Team{Id: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == uuid.MustParse("33333333-3333-3333-3333-333333333333") {
		return models.Race{Id: uuid.MustParse("33333333-3333-3333-3333-333333333333")}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpsertRaceResult(ctx context.Context, result models.RaceResult) (models.RaceResult, error) {
	if result.RaceID == uuid.MustParse("33333333-3333-3333-3333-333333333333") &&
		result.CarParticipantID == uuid.MustParse("55555555-5555-5555-5555-555555555555") {
		return result, nil
	}
	return models.RaceResult{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	if track.Name == "testName" {
		return models.Track{Id: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "testName"}, nil
	}
	return models.Track{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Number == "7" {
		return models.CarParticipant{
			Id:      uuid.MustParse("66666666-6666-6666-6666-666666666666"),
			Number:  "7",
			Drivers: carParticipant.Drivers,
		}, nil
	}
	return models.CarParticipant{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == uuid.MustParse("33333333-3333-3333-3333-333333333333") {
		return models.Race{Id: uuid.MustParse("33333333-3333-3333-3333-333333333333")}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == uuid.MustParse("22222222-2222-2222-2222-222222222222") {
		return models.Team{Id: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	if track.Id == uuid.MustParse("44444444-4444-4444-4444-444444444444") {
		return models.Track{Id: uuid.MustParse("44444444-4444-4444-4444-444444444444"), Name: "testName"}, nil
	}
	return models.Track{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.MustParse("66666666-6666-6666-6666-666666666666") {
		return models.CarParticipant{Id: uuid.MustParse("66666666-6666-6666-6666-666666666666"), Number: "7"}, nil
	}
	return models.CarParticipant{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateCarParticipantDrivers(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.MustParse("66666666-6666-6666-6666-666666666666") {
		return carParticipant, nil
	}
	return models.CarParticipant{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == uuid.MustParse("11111111-1111-1111-1111-111111111111") {
		return models.Driver{Id: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateUserRole(ctx context.Context, user models.User) (models.User, error) {
	if user.Id == uuid.MustParse("77777777-7777-7777-7777-777777777777") {
		return user, nil
	}
	return models.User{}, shared.ErrorNotFound
}
