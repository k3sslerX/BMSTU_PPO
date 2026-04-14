package bootstrap

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

type testRepo struct {
	lastUser models.User
	err      error
	calls    int
	hasAdmin bool
	hasErr   error
	checked  int
}

func (r *testRepo) HasAdmin(_ context.Context) (bool, error) {
	r.checked++
	if r.hasErr != nil {
		return false, r.hasErr
	}

	return r.hasAdmin, nil
}

func (r *testRepo) UserLogin(_ context.Context, _ models.User) (models.User, error) {
	return models.User{}, nil
}

func (r *testRepo) UserRegister(_ context.Context, user models.User) (models.User, error) {
	r.calls++
	r.lastUser = user
	if r.err != nil {
		return models.User{}, r.err
	}

	return user, nil
}

func (r *testRepo) UserChangePassword(_ context.Context, _ models.User, _ string) error {
	return nil
}

func TestEnsureAdminCreatesDefaultAdmin(t *testing.T) {
	repo := &testRepo{}

	err := EnsureAdmin(context.Background(), repo)
	if err != nil {
		t.Fatalf("EnsureAdmin() error = %v", err)
	}

	if repo.calls != 1 {
		t.Fatalf("UserRegister() calls = %d, want 1", repo.calls)
	}
	if repo.checked != 1 {
		t.Fatalf("HasAdmin() calls = %d, want 1", repo.checked)
	}
	if repo.lastUser.Name != defaultAdminName {
		t.Fatalf("user name = %q, want %q", repo.lastUser.Name, defaultAdminName)
	}
	if repo.lastUser.Email != defaultAdminLogin {
		t.Fatalf("user email = %q, want %q", repo.lastUser.Email, defaultAdminLogin)
	}
	if repo.lastUser.Password == defaultAdminPassword {
		t.Fatalf("user password was not hashed")
	}
	if repo.lastUser.Role != models.RoleAdmin {
		t.Fatalf("user role = %q, want %q", repo.lastUser.Role, models.RoleAdmin)
	}
}

func TestEnsureAdminSkipsCreationWhenAdminExists(t *testing.T) {
	repo := &testRepo{hasAdmin: true}

	err := EnsureAdmin(context.Background(), repo)
	if err != nil {
		t.Fatalf("EnsureAdmin() error = %v, want nil", err)
	}
	if repo.calls != 0 {
		t.Fatalf("UserRegister() calls = %d, want 0", repo.calls)
	}
}

func TestEnsureAdminIgnoresExistingAdminLoginCollision(t *testing.T) {
	repo := &testRepo{err: shared.ErrorUserAlreadyExists}

	err := EnsureAdmin(context.Background(), repo)
	if err != nil {
		t.Fatalf("EnsureAdmin() error = %v, want nil", err)
	}
}

func TestEnsureAdminReturnsHasAdminError(t *testing.T) {
	wantErr := errors.New("db is down")
	repo := &testRepo{hasErr: wantErr}

	err := EnsureAdmin(context.Background(), repo)
	if !errors.Is(err, wantErr) {
		t.Fatalf("EnsureAdmin() error = %v, want %v", err, wantErr)
	}
}

func TestEnsureAdminReturnsUnexpectedRegisterError(t *testing.T) {
	wantErr := errors.New("register failed")
	repo := &testRepo{err: wantErr}

	err := EnsureAdmin(context.Background(), repo)
	if !errors.Is(err, wantErr) {
		t.Fatalf("EnsureAdmin() error = %v, want %v", err, wantErr)
	}
}
