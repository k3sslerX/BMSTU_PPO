package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testAdminTeamID = uuid.MustParse("22222222-2222-2222-2222-222222222222")
var unknownAdminTeamID = uuid.MustParse("aaaaaaaa-9999-9999-9999-999999999999")

func TestUpdateTeamUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Team{Id: testAdminTeamID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Team{Id: testAdminTeamID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateTeamUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Team{Id: unknownAdminTeamID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
