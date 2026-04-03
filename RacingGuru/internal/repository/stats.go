package repository

import (
	"RacingGuru/internal/models"
	"context"
)

func (r *Repository) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	return models.DriverStats{}, nil
}

func (r *Repository) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	return models.TeamStats{}, nil
}

func (r *Repository) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	var id int
	var birthday, nationality string
	row := r.Pool.QueryRow(ctx,
		"SELECT id, name, birthday, nationality FROM driver WHERE name like '%$1'", name)
	err := row.Scan(&id, &name, &birthday, &nationality)
	if err != nil {
		return models.Driver{}, err
	}

	return models.Driver{Id: id, Name: name, Birthday: birthday, Nationality: nationality}, nil
}

func (r *Repository) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	var id int
	var country string
	row := r.Pool.QueryRow(ctx,
		"SELECT id, name, country FROM team WHERE name like '%$1'", name)
	err := row.Scan(&id, &name, &country)
	if err != nil {
		return models.Team{}, err
	}

	return models.Team{Id: id, Name: name, Country: country}, nil
}
