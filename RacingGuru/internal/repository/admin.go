package repository

import (
	"RacingGuru/internal/models"
	"context"
)

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	return driver, nil
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	return team, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
}

func (r *Repository) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	return driver, nil
}

func (r *Repository) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	return team, nil
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
}
