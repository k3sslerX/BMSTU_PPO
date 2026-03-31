package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type ToggleFavouriteTeamUseCase struct {
	Repo Repo
	User models.User
}

func NewToggleFavouriteTeamUseCase(repo Repo, user models.User) *ToggleFavouriteTeamUseCase {
	return &ToggleFavouriteTeamUseCase{Repo: repo, User: user}
}

func (uc *ToggleFavouriteTeamUseCase) Run(ctx context.Context, team models.Team) error {
	if team.Id == "" {
		var err error
		team, err = uc.Repo.GetTeamByName(ctx, team.Name)
		if err != nil {
			return err
		}
		if team.Id == "" {
			return shared.ErrorNotFound
		}
	}
	return uc.Repo.ToggleFavouriteTeam(ctx, uc.User, team)
}
