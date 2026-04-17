package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

var testAdminDriverID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
var unknownAdminDriverID = uuid.MustParse("99999999-9999-9999-9999-999999999999")

func TestUpdateDriverUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Driver{Id: testAdminDriverID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Driver{Id: testAdminDriverID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Driver{Id: unknownAdminDriverID})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
