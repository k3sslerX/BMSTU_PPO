package stats

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
)

type DriverStatsUseCase struct {
	Repo Repo
	User models.User
}

func NewDriverStatsUseCase(repo Repo, user models.User) *DriverStatsUseCase {
	return &DriverStatsUseCase{Repo: repo, User: user}
}

func (uc *DriverStatsUseCase) Run(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	if driver.Id <= 0 {
		var err error
		driver, err = uc.Repo.GetDriverByName(ctx, driver.Name)
		if err != nil {
			return models.DriverStats{}, err
		}
		if driver.Id <= 0 {
			return models.DriverStats{}, shared.ErrorNotFound
		}
	}
	return uc.Repo.GetDriverStats(ctx, driver)
}
