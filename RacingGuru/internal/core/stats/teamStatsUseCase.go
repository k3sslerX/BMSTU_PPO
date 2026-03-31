package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type TeamStatsUseCase struct {
	Repo Repo
	User models.User
}

func NewTeamStatsUseCase(repo Repo, user models.User) *TeamStatsUseCase {
	return &TeamStatsUseCase{Repo: repo, User: user}
}

func (uc *TeamStatsUseCase) Run(ctx context.Context, team models.Team) (models.TeamStats, error) {
	if team.Id == "" {
		var err error
		team, err = uc.Repo.GetTeamByName(ctx, team.Name)
		if err != nil {
			return models.TeamStats{}, err
		}
		if team.Id == "" {
			return models.TeamStats{}, shared.ErrorNotFound
		}
	}
	return uc.Repo.GetTeamStats(ctx, team)
}
