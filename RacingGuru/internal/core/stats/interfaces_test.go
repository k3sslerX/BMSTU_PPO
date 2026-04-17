package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type testRepo struct {
}

func (repo *testRepo) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	if driver.Id == uuid.MustParse("11111111-1111-1111-1111-111111111111") {
		return models.DriverStats{}, nil
	}
	return models.DriverStats{}, shared.ErrorNotFound
}

func (repo *testRepo) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	if team.Id == uuid.MustParse("22222222-2222-2222-2222-222222222222") {
		return models.TeamStats{}, nil
	}
	return models.TeamStats{}, shared.ErrorNotFound
}

func (repo *testRepo) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	if name == "testName" {
		return models.Driver{Id: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	if name == "testName" {
		return models.Team{Id: uuid.MustParse("22222222-2222-2222-2222-222222222222"), Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}
