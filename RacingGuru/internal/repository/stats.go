package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) ListDrivers(ctx context.Context, query string) ([]models.Driver, error) {
	builder := statementBuilder().
		Select("id", "name", "birthday", "nationality").
		From("driver").
		OrderBy("name ASC")

	if query != "" {
		builder = builder.Where(sq.ILike{"name": "%" + query + "%"})
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	drivers := make([]models.Driver, 0)
	for rows.Next() {
		var driver models.Driver
		var birthday time.Time
		if err := rows.Scan(&driver.Id, &driver.Name, &birthday, &driver.Nationality); err != nil {
			return nil, err
		}
		driver.Birthday = birthday.Format(time.DateOnly)
		drivers = append(drivers, driver)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

func (r *Repository) ListTeams(ctx context.Context, query string) ([]models.Team, error) {
	builder := statementBuilder().
		Select("id", "name", "country").
		From("team").
		OrderBy("name ASC")

	if query != "" {
		builder = builder.Where(sq.ILike{"name": "%" + query + "%"})
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make([]models.Team, 0)
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.Id, &team.Name, &team.Country); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *Repository) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	driverStats := models.DriverStats{}
	driverStats.Driver = driver
	stats := models.Stats{}
	var bestFinish, bestQualifying, bestChampionshipPosition sql.NullInt64
	var totalRaces, totalWins, totalPodiums, totalPoints, totalPoles, championshipsWins int64
	query, args, err := sq.Expr(
		"SELECT total_races, total_wins, total_podiums, total_points, total_poles, "+
			"best_finish, best_qualifying, championship_wins, best_championship_position "+
			"FROM CalculateDriverStats(?)",
		driver.Id,
	).ToSql()
	if err != nil {
		return driverStats, err
	}
	query, err = sq.Dollar.ReplacePlaceholders(query)
	if err != nil {
		return driverStats, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&totalRaces, &totalWins, &totalPodiums, &totalPoints, &totalPoles,
		&bestFinish, &bestQualifying, &championshipsWins, &bestChampionshipPosition)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return driverStats, shared.ErrorNotFound
		}
		return driverStats, err
	}
	stats.TotalRaces = int(totalRaces)
	stats.TotalWins = int(totalWins)
	stats.TotalPodiums = int(totalPodiums)
	stats.TotalPoints = int(totalPoints)
	stats.TotalPoles = int(totalPoles)
	stats.BestFinish = nullableInt(bestFinish)
	stats.BestQualifying = nullableInt(bestQualifying)
	stats.ChampionshipsWins = int(championshipsWins)
	stats.BestChampionshipPosition = nullableInt(bestChampionshipPosition)
	driverStats.Stats = stats

	return driverStats, nil
}

func (r *Repository) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	teamStats := models.TeamStats{}
	teamStats.Team = team
	stats := models.Stats{}
	var bestFinish, bestQualifying, bestChampionshipPosition sql.NullInt64
	var totalRaces, totalWins, totalPodiums, totalPoints, totalPoles, championshipsWins int64
	query, args, err := sq.Expr(
		"SELECT total_races, total_wins, total_podiums, total_points, total_poles, "+
			"best_finish, best_qualifying, championship_wins, best_championship_position "+
			"FROM CalculateTeamStats(?)",
		team.Id,
	).ToSql()
	if err != nil {
		return teamStats, err
	}
	query, err = sq.Dollar.ReplacePlaceholders(query)
	if err != nil {
		return teamStats, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&totalRaces, &totalWins, &totalPodiums, &totalPoints, &totalPoles,
		&bestFinish, &bestQualifying, &championshipsWins, &bestChampionshipPosition)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return teamStats, shared.ErrorNotFound
		}
		return teamStats, err
	}
	stats.TotalRaces = int(totalRaces)
	stats.TotalWins = int(totalWins)
	stats.TotalPodiums = int(totalPodiums)
	stats.TotalPoints = int(totalPoints)
	stats.TotalPoles = int(totalPoles)
	stats.BestFinish = nullableInt(bestFinish)
	stats.BestQualifying = nullableInt(bestQualifying)
	stats.ChampionshipsWins = int(championshipsWins)
	stats.BestChampionshipPosition = nullableInt(bestChampionshipPosition)
	teamStats.Stats = stats

	return teamStats, nil
}

func (r *Repository) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	var id uuid.UUID
	var nationality string
	var birthday time.Time
	query, args, err := statementBuilder().
		Select("id", "name", "birthday", "nationality").
		From("driver").
		Where(sq.Like{"name": name}).
		ToSql()
	if err != nil {
		return models.Driver{}, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&id, &name, &birthday, &nationality)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Driver{}, shared.ErrorNotFound
		}
		return models.Driver{}, err
	}

	return models.Driver{Id: id, Name: name, Birthday: birthday.Format(time.DateOnly), Nationality: nationality}, nil
}

func (r *Repository) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	var id uuid.UUID
	var country string
	query, args, err := statementBuilder().
		Select("id", "name", "country").
		From("team").
		Where(sq.Like{"name": name}).
		ToSql()
	if err != nil {
		return models.Team{}, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&id, &name, &country)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Team{}, shared.ErrorNotFound
		}
		return models.Team{}, err
	}

	return models.Team{Id: id, Name: name, Country: country}, nil
}

func nullableInt(value sql.NullInt64) int {
	if !value.Valid {
		return 0
	}

	return int(value.Int64)
}
