package auth

import (
	"RacingGuru/internal/models"
	"context"
	"testing"
)

func TestUserMeUseCase(t *testing.T) {
	uc := NewUserMeUseCase(&testRepo{}, models.User{Id: testUserID})

	user, err := uc.Run(context.Background())
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}
	if user.Name != "testName" || user.Email != "test@example.com" {
		t.Fatalf("Run() user = %+v, want test profile", user)
	}
}
