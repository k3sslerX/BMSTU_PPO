package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateDriverUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Driver{Id: testDriverID})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Driver{Id: testDriverID})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewUpdateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Driver{Id: uuid.MustParse("33333333-3333-3333-3333-333333333333")})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
