package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateCarParticipantDriversUseCase(t *testing.T) {
	repo := &testRepo{}
	input := models.CarParticipant{
		Id:      testCarParticipantAdminID,
		Drivers: []models.Driver{{Id: uuid.MustParse("77777777-7777-7777-7777-777777777777")}},
	}

	testUc1 := NewUpdateCarParticipantDriversUseCase(repo, models.User{Role: models.RoleUser})
	_, err := testUc1.Run(context.Background(), input)
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateCarParticipantDriversUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), input)
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateCarParticipantDriversUseCase(repo, models.User{Role: models.RoleAdmin})
	input.Id = unknownCarParticipantAdminID
	_, err = testUc3.Run(context.Background(), input)
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
