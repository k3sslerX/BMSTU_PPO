package auth

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

var testUserID = uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")

type testRepo struct {
	hasAdmin bool
}

func (repo *testRepo) HasAdmin(ctx context.Context) (bool, error) {
	return repo.hasAdmin, nil
}

func (repo *testRepo) GetUserByID(ctx context.Context, user models.User) (models.User, error) {
	if user.Id == testUserID {
		return models.User{Id: testUserID, Name: "testName", Email: "test@example.com", Role: models.RoleUser}, nil
	}
	return models.User{}, shared.ErrorNotFound
}

func (repo *testRepo) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "testName" && user.Password == hashPassword("testPassword") {
		return models.User{Id: testUserID, Name: "testName", Role: models.RoleUser}, nil
	}
	return models.User{}, shared.ErrorIncorrectPassword
}

func (repo *testRepo) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	if user.Name == "testName" && user.Password == hashPassword("testPassword") {
		return models.User{Id: testUserID, Name: user.Name, Email: user.Email, Role: user.Role}, nil
	}
	return models.User{}, shared.ErrorNotFound
}

func (repo *testRepo) UserChangePassword(ctx context.Context, user models.User, pwd string) error {
	if user.Id == testUserID && pwd == hashPassword("newPassword") {
		return nil
	}
	return shared.ErrorNotFound
}
