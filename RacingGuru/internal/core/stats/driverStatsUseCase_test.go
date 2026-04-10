package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testDriverID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
var unknownDriverID = uuid.MustParse("99999999-9999-9999-9999-999999999999")

func TestDriverStatsUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewDriverStatsUseCase(repo)
	_, err = testUc1.Run(context.Background(), models.Driver{})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc2 := NewDriverStatsUseCase(repo)
	_, err = testUc2.Run(context.Background(), models.Driver{Id: testDriverID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewDriverStatsUseCase(repo)
	_, err = testUc3.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}

	testUc4 := NewDriverStatsUseCase(repo)
	_, err = testUc4.Run(context.Background(), models.Driver{Id: unknownDriverID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc5 := NewDriverStatsUseCase(repo)
	_, err = testUc5.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
