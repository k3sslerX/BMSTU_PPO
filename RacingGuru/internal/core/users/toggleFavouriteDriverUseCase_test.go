package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestToggleFavouriteDriverUseCase(t *testing.T) {
	var err error
	userID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	driverID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	user := models.User{Id: userID, Role: models.RoleUser}

	repo1 := &testRepo{}
	testUc1 := NewToggleFavouriteDriverUseCase(repo1, user)
	err = testUc1.Run(context.Background(), models.Driver{Id: driverID})
	if err != nil {
		t.Error(err)
	}
	if !repo1.hasFavouriteDriver(userID, driverID) {
		t.Error("expected favourite_driver row to be inserted")
	}

	repo2 := &testRepo{favouriteDrivers: []favouriteDriverRow{{userID: userID, driverID: driverID}}}
	testUc2 := NewToggleFavouriteDriverUseCase(repo2, user)
	err = testUc2.Run(context.Background(), models.Driver{Id: driverID})
	if err != nil {
		t.Error(err)
	}
	if repo2.hasFavouriteDriver(userID, driverID) {
		t.Error("expected favourite_driver row to be deleted")
	}

	repo3 := &testRepo{}
	testUc3 := NewToggleFavouriteDriverUseCase(repo3, user)
	err = testUc3.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}
	if !repo3.hasFavouriteDriver(userID, driverID) {
		t.Error("expected resolved driver row to be inserted")
	}

	repo4 := &testRepo{}
	testUc4 := NewToggleFavouriteDriverUseCase(repo4, user)
	err = testUc4.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
