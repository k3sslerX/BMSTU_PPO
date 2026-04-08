package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestCreateTrackUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewCreateTrackUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Track{Name: "testName"})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewCreateTrackUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Track{Name: "testName"})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewCreateTrackUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Track{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
