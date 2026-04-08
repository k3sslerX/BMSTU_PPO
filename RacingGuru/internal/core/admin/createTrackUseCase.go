package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CreateTrackUseCase struct {
	Repo Repo
	User models.User
}

func NewCreateTrackUseCase(repo Repo, user models.User) *CreateTrackUseCase {
	return &CreateTrackUseCase{Repo: repo, User: user}
}

func (uc *CreateTrackUseCase) Run(ctx context.Context, track models.Track) (models.Track, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Track{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.CreateTrack(ctx, track)
}
