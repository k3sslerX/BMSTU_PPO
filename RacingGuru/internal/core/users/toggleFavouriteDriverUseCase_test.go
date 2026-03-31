package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestToggleFavouriteDriverUseCase(t *testing.T) {
	var err error
	user := models.User{Id: "userId", Role: models.RoleUser}

	repo1 := &testRepo{}
	testUc1 := NewToggleFavouriteDriverUseCase(repo1, user)
	err = testUc1.Run(context.Background(), models.Driver{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if !repo1.favouriteDrivers["testId"] || repo1.lastDriverAction != "added" {
		t.Error("expected driver to be added to favourites")
	}

	repo2 := &testRepo{favouriteDrivers: map[string]bool{"testId": true}}
	testUc2 := NewToggleFavouriteDriverUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Driver{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if repo2.favouriteDrivers["testId"] || repo2.lastDriverAction != "removed" {
		t.Error("expected driver to be removed from favourites")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteDriverUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.favouriteDrivers["testId"] || repo3.lastDriverAction != "added" {
		t.Error("expected named driver to be resolved and added to favourites")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteDriverUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
