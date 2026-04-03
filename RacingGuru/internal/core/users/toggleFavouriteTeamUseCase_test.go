package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestToggleFavouriteTeamUseCase(t *testing.T) {
	var err error
	userID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	teamID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	user := models.User{Id: userID, Role: models.RoleUser}

	repo1 := &testRepo{}
	testUc1 := NewToggleFavouriteTeamUseCase(repo1, user)
	err = testUc1.Run(context.Background(), models.Team{Id: teamID})
	if err != nil {
		t.Error(err)
	}
	if !repo1.hasFavouriteTeam(userID, teamID) {
		t.Error("expected favourite_team row to be inserted")
	}

	repo2 := &testRepo{favouriteTeams: []favouriteTeamRow{{userID: userID, teamID: teamID}}}
	testUc2 := NewToggleFavouriteTeamUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Team{Id: teamID})
	if err != nil {
		t.Error(err)
	}
	if repo2.hasFavouriteTeam(userID, teamID) {
		t.Error("expected favourite_team row to be deleted")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteTeamUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Team{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.hasFavouriteTeam(userID, teamID) {
		t.Error("expected resolved team row to be inserted")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteTeamUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Team{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
