package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestCreateRaceUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Race{Id: 1})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Race{Id: 1})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Race{Id: 2})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
