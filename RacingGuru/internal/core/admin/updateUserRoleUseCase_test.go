package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testManagedUserID = uuid.MustParse("77777777-7777-7777-7777-777777777777")
var unknownManagedUserID = uuid.MustParse("88888888-8888-8888-8888-888888888888")

func TestUpdateUserRoleUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateUserRoleUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.User{Id: testManagedUserID, Role: models.RoleAdmin})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateUserRoleUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.User{Id: testManagedUserID, Role: models.RoleAdmin})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateUserRoleUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.User{Id: unknownManagedUserID, Role: models.RoleAdmin})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}

	testUc4 := NewUpdateUserRoleUseCase(repo, models.User{Id: testManagedUserID, Role: models.RoleAdmin})
	_, err = testUc4.Run(context.Background(), models.User{Id: testManagedUserID, Role: models.RoleUser})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}
}
