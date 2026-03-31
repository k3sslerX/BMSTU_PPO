package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestUpdateTeamUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Team{Id: "testId"})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Team{Id: "testId"})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Team{Id: "someId"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
