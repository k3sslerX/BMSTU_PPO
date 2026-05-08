package users

import (
	"RacingGuru/internal/models"
	"context"
)

type ListFavouritesUseCase struct {
	Repo Repo
	User models.User
}

func NewListFavouritesUseCase(repo Repo, user models.User) *ListFavouritesUseCase {
	return &ListFavouritesUseCase{Repo: repo, User: user}
}

func (uc *ListFavouritesUseCase) Run(ctx context.Context) ([]models.Driver, []models.Team, error) {
	drivers, err := uc.Repo.ListFavouriteDrivers(ctx, uc.User)
	if err != nil {
		return nil, nil, err
	}

	teams, err := uc.Repo.ListFavouriteTeams(ctx, uc.User)
	if err != nil {
		return nil, nil, err
	}

	return drivers, teams, nil
}
