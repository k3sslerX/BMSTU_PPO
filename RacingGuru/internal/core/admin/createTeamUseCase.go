package admin

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type CreateTeamUseCase struct {
	Repo Repo
	User models.User
}

func NewCreateTeamUseCase(repo Repo, user models.User) *CreateTeamUseCase {
	return &CreateTeamUseCase{Repo: repo, User: user}
}

func (uc *CreateTeamUseCase) Run(ctx context.Context, team models.Team) (models.Team, error) {
	if uc.User.Role != models.RoleAdmin {
		return models.Team{}, shared.ErrorPermissionDenied
	}
	return uc.Repo.CreateTeam(ctx, team)
}
