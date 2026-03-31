package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
}

func (repo *testRepo) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Name == "testName" {
		return models.Driver{Id: "testId", Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Name == "testName" {
		return models.Team{Id: "testId", Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == "testId" {
		return models.Race{Id: "testId"}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == "testId" {
		return models.Race{Id: "testId"}, nil
	}
	return models.Race{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == "testId" {
		return models.Team{Id: "testId", Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}

func (repo *testRepo) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == "testId" {
		return models.Driver{Id: "testId", Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}
