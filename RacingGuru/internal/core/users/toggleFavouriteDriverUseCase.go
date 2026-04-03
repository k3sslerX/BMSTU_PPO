package users

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"

	"github.com/google/uuid"
)

type ToggleFavouriteDriverUseCase struct {
	Repo Repo
	User models.User
}

func NewToggleFavouriteDriverUseCase(repo Repo, user models.User) *ToggleFavouriteDriverUseCase {
	return &ToggleFavouriteDriverUseCase{Repo: repo, User: user}
}

func (uc *ToggleFavouriteDriverUseCase) Run(ctx context.Context, driver models.Driver) error {
	if driver.Id == uuid.Nil {
		var err error
		driver, err = uc.Repo.GetDriverByName(ctx, driver.Name)
		if err != nil {
			return err
		}
		if driver.Id == uuid.Nil {
			return shared.ErrorNotFound
		}
	}
	return uc.Repo.ToggleFavouriteDriver(ctx, uc.User, driver)
}
