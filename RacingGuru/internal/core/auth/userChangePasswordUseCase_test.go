package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestUserChangePasswordUseCase(t *testing.T) {
	repo := &testRepo{}

	testUc1 := NewUserChangePasswordUseCase(repo, models.User{Id: testUserID})
	err := testUc1.Run(context.Background(), "newPassword")
	if err != nil {
		t.Error(err)
	}

	testUc2 := NewUserChangePasswordUseCase(repo, models.User{})
	err = testUc2.Run(context.Background(), "somePassword")
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
