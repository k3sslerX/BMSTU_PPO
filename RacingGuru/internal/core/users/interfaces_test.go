package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
	favouriteDrivers map[string]bool
	favouriteTeams   map[string]bool
	lastDriverAction string
	lastTeamAction   string
}

func (repo *testRepo) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	if repo.favouriteDrivers == nil {
		repo.favouriteDrivers = make(map[string]bool)
	}
	if repo.favouriteDrivers[driver.Id] {
		delete(repo.favouriteDrivers, driver.Id)
		repo.lastDriverAction = "removed"
		return nil
	}
	repo.favouriteDrivers[driver.Id] = true
	repo.lastDriverAction = "added"
	return nil
}

func (repo *testRepo) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	if repo.favouriteTeams == nil {
		repo.favouriteTeams = make(map[string]bool)
	}
	if repo.favouriteTeams[team.Id] {
		delete(repo.favouriteTeams, team.Id)
		repo.lastTeamAction = "removed"
		return nil
	}
	repo.favouriteTeams[team.Id] = true
	repo.lastTeamAction = "added"
	return nil
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
