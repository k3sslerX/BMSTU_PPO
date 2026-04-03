package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateTeamUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Team{Id: testTeamID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Team{Id: testTeamID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Team{Id: uuid.MustParse("44444444-4444-4444-4444-444444444444")})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
