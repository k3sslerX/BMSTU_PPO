package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestCreateDriverUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewCreateDriverUseCase(repo, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background(), models.Driver{Name: "testName"})
	if !errors.Is(err, shared.ErrorPermissionDenied) {
		t.Error(err)
	}

	testUc2 := NewCreateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc2.Run(context.Background(), models.Driver{Name: "testName"})
	if err != nil {
		t.Error(err)
	}

	testUc3 := NewCreateDriverUseCase(repo, models.User{Role: models.RoleAdmin})
	_, err = testUc3.Run(context.Background(), models.Driver{Name: "someName"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
