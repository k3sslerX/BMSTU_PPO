package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestDriverStatsUseCase(t *testing.T) {
	repo := &testRepo{}
	user := models.User{Role: models.RoleUser}
	var err error

	testUc1 := NewDriverStatsUseCase(repo, user)
	_, err = testUc1.Run(context.Background(), models.Driver{})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc2 := NewDriverStatsUseCase(repo, user)
	_, err = testUc2.Run(context.Background(), models.Driver{Id: "testId"})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewDriverStatsUseCase(repo, user)
	_, err = testUc3.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}

	testUc4 := NewDriverStatsUseCase(repo, user)
	_, err = testUc4.Run(context.Background(), models.Driver{Id: "someId"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc5 := NewDriverStatsUseCase(repo, user)
	_, err = testUc5.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
