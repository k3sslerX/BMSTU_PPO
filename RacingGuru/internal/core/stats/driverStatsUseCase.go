package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type DriverStatsUseCase struct {
	Repo Repo
}

func NewDriverStatsUseCase(repo Repo) *DriverStatsUseCase {
	return &DriverStatsUseCase{Repo: repo}
}

func (uc *DriverStatsUseCase) Run(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	if driver.Id == uuid.Nil {
		var err error
		driver, err = uc.Repo.GetDriverByName(ctx, driver.Name)
		if err != nil {
			return models.DriverStats{}, err
		}
		if driver.Id == uuid.Nil {
			return models.DriverStats{}, shared.ErrorNotFound
		}
	}
	return uc.Repo.GetDriverStats(ctx, driver)
}
