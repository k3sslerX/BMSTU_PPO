package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestUserRegisterUseCase(t *testing.T) {
	repo := &testRepo{}
	var err error

	testUc1 := NewUserRegisterUseCase(repo)
	_, err = testUc1.Run(context.Background(), models.User{Name: "testName", Password: "testPassword"})
	if err != nil {
		t.Error(err)
	}

	testUc2 := NewUserRegisterUseCase(repo)
	_, err = testUc2.Run(context.Background(), models.User{Name: "someName", Password: "somePassword"})
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
