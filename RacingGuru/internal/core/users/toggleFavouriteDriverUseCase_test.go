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
	if !repo1.hasFavouriteDriver(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected favourite_driver row to be inserted")
	}

	repo2 := &testRepo{favouriteDrivers: []favouriteDriverRow{{userID: models.Uuid("userId"), driverID: models.Uuid("testId")}}}
	testUc2 := NewToggleFavouriteDriverUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Driver{Id: "testId"})
	if err != nil {
		t.Error(err)
	}
	if repo2.hasFavouriteDriver(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected favourite_driver row to be deleted")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteDriverUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.hasFavouriteDriver(models.Uuid("userId"), models.Uuid("testId")) {
		t.Error("expected resolved driver row to be inserted")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteDriverUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
