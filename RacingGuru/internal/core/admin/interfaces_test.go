package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type testRepo struct {
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

func (repo *testRepo) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == uuid.MustParse("11111111-1111-1111-1111-111111111111") {
		return models.Driver{Id: uuid.MustParse("11111111-1111-1111-1111-111111111111"), Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}
