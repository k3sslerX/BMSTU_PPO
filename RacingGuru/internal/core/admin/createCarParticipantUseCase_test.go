package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestCreateCarParticipantUseCase(t *testing.T) {
	repo := &testRepo{}
	input := models.CarParticipant{Number: "7"}

	testUc1 := NewCreateCarParticipantUseCase(repo, models.User{Role: models.RoleUser})
	_, err := testUc1.Run(context.Background(), input)
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewCreateCarParticipantUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), input)
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewCreateCarParticipantUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.CarParticipant{Number: "404"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
