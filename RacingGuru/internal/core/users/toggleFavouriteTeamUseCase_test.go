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
	if !repo1.favouriteTeams["testId"] || repo1.lastTeamAction != "added" {
		t.Error("expected team to be added to favourites")
	}

	repo2 := &testRepo{favouriteTeams: map[string]bool{"testId": true}}
	testUc2 := NewToggleFavouriteTeamUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Team{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if repo2.favouriteTeams["testId"] || repo2.lastTeamAction != "removed" {
		t.Error("expected team to be removed from favourites")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteTeamUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Team{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.favouriteTeams["testId"] || repo3.lastTeamAction != "added" {
		t.Error("expected named team to be resolved and added to favourites")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteTeamUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Team{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
