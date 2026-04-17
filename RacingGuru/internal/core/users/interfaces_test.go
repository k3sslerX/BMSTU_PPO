package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type testRepo struct {
	favouriteDrivers []favouriteDriverRow
	favouriteTeams   []favouriteTeamRow
}

type favouriteDriverRow struct {
	userID   uuid.UUID
	driverID uuid.UUID
}

type favouriteTeamRow struct {
	userID uuid.UUID
	teamID uuid.UUID
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

func (repo *testRepo) hasFavouriteDriver(userID uuid.UUID, driverID uuid.UUID) bool {
	for _, row := range repo.favouriteDrivers {
		if row.userID == userID && row.driverID == driverID {
			return true
		}
	}
	return false
}

func (repo *testRepo) hasFavouriteTeam(userID uuid.UUID, teamID uuid.UUID) bool {
	for _, row := range repo.favouriteTeams {
		if row.userID == userID && row.teamID == teamID {
			return true
		}
	}
	return false
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
