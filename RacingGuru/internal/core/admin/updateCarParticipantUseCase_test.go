package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testCarParticipantAdminID = uuid.MustParse("66666666-6666-6666-6666-666666666666")
var unknownCarParticipantAdminID = uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")

func TestUpdateCarParticipantUseCase(t *testing.T) {
	repo := &testRepo{}

	testUc1 := NewUpdateCarParticipantUseCase(repo, models.User{Role: models.RoleUser})
	_, err := testUc1.Run(context.Background(), models.CarParticipant{Id: testCarParticipantAdminID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateCarParticipantUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.CarParticipant{Id: testCarParticipantAdminID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateCarParticipantUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.CarParticipant{Id: unknownCarParticipantAdminID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
