package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type UpdateTeamUseCase struct {
	Repo Repo
	User models.User
}

func NewUpdateTeamUseCase(repo Repo, user models.User) *UpdateTeamUseCase {
	return &UpdateTeamUseCase{Repo: repo, User: user}
}

func (uc *UpdateTeamUseCase) Run(ctx context.Context, team models.Team) (models.Team, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Team{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.UpdateTeam(ctx, team)
}
