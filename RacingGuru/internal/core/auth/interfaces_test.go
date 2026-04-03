package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

var testUserID = uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

type testRepo struct {
}

func (repo *testRepo) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "testName" && user.Password == "testPassword" {
		return models.User{Id: testUserID, Name: "testName", Role: models.RoleUser}, nil
	}
	return models.User{}, shared.ErrorNotFound
}

func (repo *testRepo) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "testName" && user.Password == "testPassword" {
		return models.User{Id: testUserID, Name: "testName", Role: models.RoleUser}, nil
	}
	return models.User{}, shared.ErrorNotFound
}
