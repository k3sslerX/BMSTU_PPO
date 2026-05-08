package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type DeleteDriverUseCase struct {
	Repo Repo
	User models.User
}

func NewDeleteDriverUseCase(repo Repo, user models.User) *DeleteDriverUseCase {
	return &DeleteDriverUseCase{Repo: repo, User: user}
}

func (uc *DeleteDriverUseCase) Run(ctx context.Context, id uuid.UUID) error {
	if uc.User.Role != models.RoleAdmin {
		return shared.ErrorPermissionDenied
	}
	return uc.Repo.DeleteDriver(ctx, id)
}

type DeleteTeamUseCase struct {
	Repo Repo
	User models.User
}

func NewDeleteTeamUseCase(repo Repo, user models.User) *DeleteTeamUseCase {
	return &DeleteTeamUseCase{Repo: repo, User: user}
}

func (uc *DeleteTeamUseCase) Run(ctx context.Context, id uuid.UUID) error {
	if uc.User.Role != models.RoleAdmin {
		return shared.ErrorPermissionDenied
	}
	return uc.Repo.DeleteTeam(ctx, id)
}

type DeleteTrackUseCase struct {
	Repo Repo
	User models.User
}

func NewDeleteTrackUseCase(repo Repo, user models.User) *DeleteTrackUseCase {
	return &DeleteTrackUseCase{Repo: repo, User: user}
}

func (uc *DeleteTrackUseCase) Run(ctx context.Context, id uuid.UUID) error {
	if uc.User.Role != models.RoleAdmin {
		return shared.ErrorPermissionDenied
	}
	return uc.Repo.DeleteTrack(ctx, id)
}

type DeleteRaceUseCase struct {
	Repo Repo
	User models.User
}

func NewDeleteRaceUseCase(repo Repo, user models.User) *DeleteRaceUseCase {
	return &DeleteRaceUseCase{Repo: repo, User: user}
}

func (uc *DeleteRaceUseCase) Run(ctx context.Context, id uuid.UUID) error {
	if uc.User.Role != models.RoleAdmin {
		return shared.ErrorPermissionDenied
	}
	return uc.Repo.DeleteRace(ctx, id)
}

type DeleteCarParticipantUseCase struct {
	Repo Repo
	User models.User
}

func NewDeleteCarParticipantUseCase(repo Repo, user models.User) *DeleteCarParticipantUseCase {
	return &DeleteCarParticipantUseCase{Repo: repo, User: user}
}

func (uc *DeleteCarParticipantUseCase) Run(ctx context.Context, id uuid.UUID) error {
	if uc.User.Role != models.RoleAdmin {
		return shared.ErrorPermissionDenied
	}
	return uc.Repo.DeleteCarParticipant(ctx, id)
}
