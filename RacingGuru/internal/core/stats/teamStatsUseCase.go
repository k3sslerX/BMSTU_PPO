package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type TeamStatsUseCase struct {
	Repo Repo
}

func NewTeamStatsUseCase(repo Repo) *TeamStatsUseCase {
	return &TeamStatsUseCase{Repo: repo}
}

func (uc *TeamStatsUseCase) Run(ctx context.Context, team models.Team) (models.TeamStats, error) {
	if team.Id == uuid.Nil {
		var err error
		team, err = uc.Repo.GetTeamByName(ctx, team.Name)
		if err != nil {
			return models.TeamStats{}, err
		}
		if team.Id == uuid.Nil {
			return models.TeamStats{}, shared.ErrorNotFound
		}
	}
	return uc.Repo.GetTeamStats(ctx, team)
}
