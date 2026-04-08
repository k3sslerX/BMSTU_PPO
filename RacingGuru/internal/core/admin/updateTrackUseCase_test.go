package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testAdminTrackID = uuid.MustParse("44444444-4444-4444-4444-444444444444")
var unknownAdminTrackID = uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")

func TestUpdateTrackUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateTrackUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Track{Id: testAdminTrackID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateTrackUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Track{Id: testAdminTrackID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateTrackUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Track{Id: unknownAdminTrackID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
