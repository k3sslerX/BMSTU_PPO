package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type testRepo struct {
}

var (
	testDriverID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
	testTeamID   = uuid.MustParse("22222222-2222-2222-2222-222222222222")
	testRaceID   = uuid.MustParse("55555555-5555-5555-5555-555555555555")
)

func (repo *testRepo) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Name == "testName" {
		return models.Driver{Id: testDriverID, Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Name == "testName" {
		return models.Team{Id: testTeamID, Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == testRaceID {
		return models.Race{Id: testRaceID}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == testRaceID {
		return models.Race{Id: testRaceID}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == testTeamID {
		return models.Team{Id: testTeamID, Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == testDriverID {
		return models.Driver{Id: testDriverID, Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}
