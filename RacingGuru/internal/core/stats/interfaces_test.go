package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
}

func (repo *testRepo) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	if driver.Id == "testId" {
		return models.DriverStats{}, nil
	}
	return models.DriverStats{}, shared.ErrorNotFound
}

func (repo *testRepo) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	if team.Id == "testId" {
		return models.TeamStats{}, nil
	}
	return models.TeamStats{}, shared.ErrorNotFound
}

func (repo *testRepo) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	if name == "testName" {
		return models.Driver{Id: "testId", Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	if name == "testName" {
		return models.Team{Id: "testId", Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}
