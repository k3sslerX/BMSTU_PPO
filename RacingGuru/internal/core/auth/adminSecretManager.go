package auth

import (
	"RacingGuru/internal/shared"
	"context"
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type AdminSecretManager struct {
	mu     sync.Mutex
	secret string
}

func NewAdminSecretManager() *AdminSecretManager {
	return &AdminSecretManager{}
}

func (m *AdminSecretManager) Generate(ctx context.Context, repo Repo) (string, error) {
	hasAdmin, err := repo.HasAdmin(ctx)
	if err != nil {
		return "", err
	}
	if hasAdmin {
		return "", shared.ErrorAdminAlreadyExists
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	if m.secret != "" {
		return "", shared.ErrorAdminSecretAlreadyIssued
	}

	buf := make([]byte, 24)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	m.secret = base64.RawURLEncoding.EncodeToString(buf)
	return m.secret, nil
}

func (m *AdminSecretManager) Validate(ctx context.Context, repo Repo, secret string) error {
	hasAdmin, err := repo.HasAdmin(ctx)
	if err != nil {
		return err
	}
	if hasAdmin {
		return shared.ErrorAdminAlreadyExists
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	if m.secret == "" || secret == "" || m.secret != secret {
		return shared.ErrorInvalidAdminSecret
	}

	return nil
}

func (m *AdminSecretManager) Consume(secret string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.secret == secret {
		m.secret = ""
	}
}
