package repository

import (
	"RacingGuru/internal/models"
	"context"

	"github.com/google/uuid"
)

func (r *Repository) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	driverStats := models.DriverStats{}
	driverStats.Driver = driver
	stats := models.Stats{}
	row := r.Pool.QueryRow(ctx,
		"SELECT total_races, total_wins, total_podiums, total_points, total_poles "+
			"best_finish, best_qualifying, championship_wins, best_championship_position "+
			"FROM CalculateDriverStats($1)", driver.Id)
	err := row.Scan(&stats.TotalRaces, &stats.TotalWins, &stats.TotalPodiums, &stats.TotalPoints, &stats.TotalPoles,
		&stats.BestFinish, &stats.BestQualifying, &stats.ChampionshipsWins, &stats.BestChampionshipPosition)
	if err != nil {
		return driverStats, err
	}
	driverStats.Stats = stats

	return driverStats, nil
}

func (r *Repository) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	teamStats := models.TeamStats{}
	teamStats.Team = team
	stats := models.Stats{}
	row := r.Pool.QueryRow(ctx,
		"SELECT total_races, total_wins, total_podiums, total_points, total_poles "+
			"best_finish, best_qualifying, championship_wins, best_championship_position "+
			"FROM CalculateTeamStats($1)", team.Id)
	err := row.Scan(&stats.TotalRaces, &stats.TotalWins, &stats.TotalPodiums, &stats.TotalPoints, &stats.TotalPoles,
		&stats.BestFinish, &stats.BestQualifying, &stats.ChampionshipsWins, &stats.BestChampionshipPosition)
	if err != nil {
		return teamStats, err
	}
	teamStats.Stats = stats

	return teamStats, nil
}

func (r *Repository) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	var id uuid.UUID
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
	var id uuid.UUID
	var country string
	row := r.Pool.QueryRow(ctx,
		"SELECT id, name, country FROM team WHERE name like '%$1'", name)
	err := row.Scan(&id, &name, &country)
	if err != nil {
		return models.Team{}, err
	}

	return models.Team{Id: id, Name: name, Country: country}, nil
}
