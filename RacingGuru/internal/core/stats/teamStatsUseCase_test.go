package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestTeamStatsUseCase(t *testing.T) {
	repo := &testRepo{}
	user := models.User{Role: models.RoleUser}
	var err error

	testUc1 := NewTeamStatsUseCase(repo, user)
	_, err = testUc1.Run(context.Background(), models.Team{})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc2 := NewTeamStatsUseCase(repo, user)
	_, err = testUc2.Run(context.Background(), models.Team{Id: 1})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewTeamStatsUseCase(repo, user)
	_, err = testUc3.Run(context.Background(), models.Team{Name: "testName"})
	if err != nil {
		t.Error(err)
	}

	testUc4 := NewTeamStatsUseCase(repo, user)
	_, err = testUc4.Run(context.Background(), models.Team{Id: 2})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc5 := NewTeamStatsUseCase(repo, user)
	_, err = testUc5.Run(context.Background(), models.Team{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
