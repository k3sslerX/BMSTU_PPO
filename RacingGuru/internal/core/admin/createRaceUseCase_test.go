package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testRaceID = uuid.MustParse("33333333-3333-3333-3333-333333333333")
var unknownRaceID = uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc")

func TestCreateRaceUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Race{Id: testRaceID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Race{Id: testRaceID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewCreateRaceUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Race{Id: unknownRaceID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
