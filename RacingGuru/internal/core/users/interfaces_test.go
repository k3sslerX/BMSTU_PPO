package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type testRepo struct {
	favouriteDrivers []favouriteDriverRow
	favouriteTeams   []favouriteTeamRow
}

type favouriteDriverRow struct {
	userID   models.Uuid
	driverID int
}

type favouriteTeamRow struct {
	userID models.Uuid
	teamID int
}

func (repo *testRepo) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	for i, row := range repo.favouriteDrivers {
		if row.userID == user.Id && row.driverID == driver.Id {
			repo.favouriteDrivers = append(repo.favouriteDrivers[:i], repo.favouriteDrivers[i+1:]...)
			return nil
		}
	}
	repo.favouriteDrivers = append(repo.favouriteDrivers, favouriteDriverRow{userID: user.Id, driverID: driver.Id})
	return nil
}

func (repo *testRepo) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	for i, row := range repo.favouriteTeams {
		if row.userID == user.Id && row.teamID == team.Id {
			repo.favouriteTeams = append(repo.favouriteTeams[:i], repo.favouriteTeams[i+1:]...)
			return nil
		}
	}
	repo.favouriteTeams = append(repo.favouriteTeams, favouriteTeamRow{userID: user.Id, teamID: team.Id})
	return nil
}

func (repo *testRepo) hasFavouriteDriver(userID models.Uuid, driverID int) bool {
	for _, row := range repo.favouriteDrivers {
		if row.userID == userID && row.driverID == driverID {
			return true
		}
	}
	return false
}

func (repo *testRepo) hasFavouriteTeam(userID models.Uuid, teamID int) bool {
	for _, row := range repo.favouriteTeams {
		if row.userID == userID && row.teamID == teamID {
			return true
		}
	}
	return false
}

func (repo *testRepo) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	if name == "testName" {
		return models.Driver{Id: 1, Name: "testName"}, nil
	}
	return models.Driver{}, shared.ErrorNotFound
}

func (repo *testRepo) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	if name == "testName" {
		return models.Team{Id: 1, Name: "testName"}, nil
	}
	return models.Team{}, shared.ErrorNotFound
}
