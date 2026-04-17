package auth

import (
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestAdminSecretManagerGenerateAndValidate(t *testing.T) {
	repo := &testRepo{}
	manager := NewAdminSecretManager()

	secret, err := manager.Generate(context.Background(), repo)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if secret == "" {
		t.Fatal("Generate() returned empty secret")
	}

	if err := manager.Validate(context.Background(), repo, secret); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestAdminSecretManagerGenerateOnce(t *testing.T) {
	repo := &testRepo{}
	manager := NewAdminSecretManager()

	if _, err := manager.Generate(context.Background(), repo); err != nil {
		t.Fatalf("first Generate() error = %v", err)
	}

	_, err := manager.Generate(context.Background(), repo)
	if !errors.Is(err, shared.ErrorAdminSecretAlreadyIssued) {
		t.Fatalf("second Generate() error = %v, want %v", err, shared.ErrorAdminSecretAlreadyIssued)
	}
}

func TestAdminSecretManagerRejectsInvalidSecret(t *testing.T) {
	repo := &testRepo{}
	manager := NewAdminSecretManager()

	if _, err := manager.Generate(context.Background(), repo); err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	err := manager.Validate(context.Background(), repo, "wrong")
	if !errors.Is(err, shared.ErrorInvalidAdminSecret) {
		t.Fatalf("Validate() error = %v, want %v", err, shared.ErrorInvalidAdminSecret)
	}
}

func TestAdminSecretManagerRejectsWhenAdminExists(t *testing.T) {
	repo := &testRepo{hasAdmin: true}
	manager := NewAdminSecretManager()

	_, err := manager.Generate(context.Background(), repo)
	if !errors.Is(err, shared.ErrorAdminAlreadyExists) {
		t.Fatalf("Generate() error = %v, want %v", err, shared.ErrorAdminAlreadyExists)
	}
}
