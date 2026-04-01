package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestToggleFavouriteTeamUseCase(t *testing.T) {
	var err error
	user := models.User{Id: "userId", Role: models.RoleUser}

	repo1 := &testRepo{}
	testUc1 := NewToggleFavouriteTeamUseCase(repo1, user)
	err = testUc1.Run(context.Background(), models.Team{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if !repo1.hasFavouriteTeam(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected favourite_team row to be inserted")
	}

	repo2 := &testRepo{favouriteTeams: []favouriteTeamRow{{userID: models.Uuid("userId"), teamID: models.Uuid("testId")}}}
	testUc2 := NewToggleFavouriteTeamUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Team{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if repo2.hasFavouriteTeam(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected favourite_team row to be deleted")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteTeamUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Team{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.hasFavouriteTeam(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected resolved team row to be inserted")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteTeamUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Team{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
