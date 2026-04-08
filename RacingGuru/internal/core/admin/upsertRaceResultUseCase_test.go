package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testCarParticipantID = uuid.MustParse("55555555-5555-5555-5555-555555555555")
var unknownCarParticipantID = uuid.MustParse("eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee")

func TestUpsertRaceResultUseCase(t *testing.T) {
	repo := &testRepo{}
	input := models.RaceResult{
		RaceID:           testRaceID,
		CarParticipantID: testCarParticipantID,
		FinishPos:        1,
		QualifyingPos:    2,
	}

	testUc1 := NewUpsertRaceResultUseCase(repo, models.User{Role: models.RoleUser})
	_, err := testUc1.Run(context.Background(), input)
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpsertRaceResultUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), input)
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpsertRaceResultUseCase(repo, models.User{Role: models.RoleAdmin})
	input.CarParticipantID = unknownCarParticipantID
	_, err = testUc3.Run(context.Background(), input)
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
