package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateTrackUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateTrackUseCase(repo Repo, user models.User) *UpdateTrackUseCase {
	return &UpdateTrackUseCase{Repo: repo, User: user}
}

func (uc *UpdateTrackUseCase) Run(ctx context.Context, track models.Track) (models.Track, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Track{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateTrack(ctx, track)
}
