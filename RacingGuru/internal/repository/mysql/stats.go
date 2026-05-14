package mysql

import (
	"context"
	"database/sql"

	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) ListDrivers(ctx context.Context, query string) ([]models.Driver, error) {
	builder := statementBuilder().
		Select("id", "name", dateOnlyExpression("birthday"), "nationality").
		From("driver").
		OrderBy("name ASC")

	if query != "" {
		builder = builder.Where(sq.Like{"name": contains(query)})
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	drivers := make([]models.Driver, 0)
	for rows.Next() {
		var driver models.Driver
		var id string
		if err := rows.Scan(&id, &driver.Name, &driver.Birthday, &driver.Nationality); err != nil {
			return nil, err
		}
		driver.Id, err = parseUUID(id)
		if err != nil {
			return nil, err
		}
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
		builder = builder.Where(sq.Like{"name": contains(query)})
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make([]models.Team, 0)
	for rows.Next() {
		var team models.Team
		var id string
		if err := rows.Scan(&id, &team.Name, &team.Country); err != nil {
			return nil, err
		}
		team.Id, err = parseUUID(id)
		if err != nil {
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
	driverStats := models.DriverStats{Driver: driver}
	stats, err := r.scanStats(r.DB.QueryRowContext(ctx, driverStatsQuery, driver.Id.String()))
	if err != nil {
		return driverStats, err
	}
	driverStats.Stats = stats

	return driverStats, nil
}

func (r *Repository) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	teamStats := models.TeamStats{Team: team}
	stats, err := r.scanStats(r.DB.QueryRowContext(ctx, teamStatsQuery, team.Id.String(), team.Id.String()))
	if err != nil {
		return teamStats, err
	}
	teamStats.Stats = stats

	return teamStats, nil
}

func (r *Repository) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	query, args, err := statementBuilder().
		Select("id", "name", dateOnlyExpression("birthday"), "nationality").
		From("driver").
		Where(sq.Eq{"name": name}).
		ToSql()
	if err != nil {
		return models.Driver{}, err
	}

	var driver models.Driver
	var id string
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&id, &driver.Name, &driver.Birthday, &driver.Nationality)
	if err != nil {
		if isNoRows(err) {
			return models.Driver{}, shared.ErrorNotFound
		}
		return models.Driver{}, err
	}

	driver.Id, err = parseUUID(id)
	if err != nil {
		return models.Driver{}, err
	}

	return driver, nil
}

func (r *Repository) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	query, args, err := statementBuilder().
		Select("id", "name", "country").
		From("team").
		Where(sq.Eq{"name": name}).
		ToSql()
	if err != nil {
		return models.Team{}, err
	}

	var team models.Team
	var id string
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&id, &team.Name, &team.Country)
	if err != nil {
		if isNoRows(err) {
			return models.Team{}, shared.ErrorNotFound
		}
		return models.Team{}, err
	}

	team.Id, err = parseUUID(id)
	if err != nil {
		return models.Team{}, err
	}

	return team, nil
}

func (r *Repository) scanStats(row rowScanner) (models.Stats, error) {
	var stats models.Stats
	var totalRaces, totalWins, totalPodiums, totalPoints, totalPoles sql.NullInt64
	var bestFinish, bestQualifying, championshipsWins, bestChampionshipPosition sql.NullInt64

	err := row.Scan(
		&totalRaces,
		&totalWins,
		&totalPodiums,
		&totalPoints,
		&totalPoles,
		&bestFinish,
		&bestQualifying,
		&championshipsWins,
		&bestChampionshipPosition,
	)
	if err != nil {
		if isNoRows(err) {
			return stats, shared.ErrorNotFound
		}
		return stats, err
	}

	stats.TotalRaces = nullableInt(totalRaces)
	stats.TotalWins = nullableInt(totalWins)
	stats.TotalPodiums = nullableInt(totalPodiums)
	stats.TotalPoints = nullableInt(totalPoints)
	stats.TotalPoles = nullableInt(totalPoles)
	stats.BestFinish = nullableInt(bestFinish)
	stats.BestQualifying = nullableInt(bestQualifying)
	stats.ChampionshipsWins = nullableInt(championshipsWins)
	stats.BestChampionshipPosition = nullableInt(bestChampionshipPosition)

	return stats, nil
}

const driverStatsQuery = `
WITH driver_lineups AS (
    SELECT DISTINCT
        tp.car_p AS car_p_id,
        rc.name AS raceclass_name
    FROM team_p tp
    JOIN car_p cp ON cp.id = tp.car_p
    JOIN car c ON c.id = cp.car
    JOIN raceclass rc ON rc.id = c.raceclass
    WHERE tp.driver = ?
),
race_results AS (
    SELECT
        r.id AS race_id,
        r.championship AS championship_id,
        cp.id AS car_p_id,
        f.pos AS finish_pos,
        CEIL(
            CASE f.pos
                WHEN 1 THEN 25
                WHEN 2 THEN 18
                WHEN 3 THEN 15
                WHEN 4 THEN 12
                WHEN 5 THEN 10
                WHEN 6 THEN 8
                WHEN 7 THEN 6
                WHEN 8 THEN 4
                WHEN 9 THEN 2
                WHEN 10 THEN 1
                ELSE 0
            END *
            CASE
                WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                WHEN r.type = 2 AND r.duration >= 24 THEN 2
                ELSE 1
            END
        ) AS points
    FROM driver_lineups dl
    JOIN car_p cp ON cp.id = dl.car_p_id
    JOIN finish f ON f.car_p = cp.id
    JOIN race r ON r.id = f.race
),
qualifying_results AS (
    SELECT q.pos
    FROM qualifying q
    JOIN driver_lineups dl ON dl.car_p_id = q.car_p
),
driver_championships AS (
    SELECT DISTINCT championship_id
    FROM race_results
),
personal_points AS (
    SELECT
        r.championship AS championship_id,
        cp.id AS car_p_id,
        rc.name AS raceclass,
        SUM(
            CEIL(
                CASE f.pos
                    WHEN 1 THEN 25
                    WHEN 2 THEN 18
                    WHEN 3 THEN 15
                    WHEN 4 THEN 12
                    WHEN 5 THEN 10
                    WHEN 6 THEN 8
                    WHEN 7 THEN 6
                    WHEN 8 THEN 4
                    WHEN 9 THEN 2
                    WHEN 10 THEN 1
                    ELSE 0
                END *
                CASE
                    WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                    WHEN r.type = 2 AND r.duration >= 24 THEN 2
                    ELSE 1
                END
            )
        ) AS points
    FROM finish f
    JOIN car_p cp ON f.car_p = cp.id
    JOIN race r ON f.race = r.id
    JOIN car c ON cp.car = c.id
    JOIN raceclass rc ON c.raceclass = rc.id
    GROUP BY r.championship, cp.id, rc.name
),
championship_positions AS (
    SELECT
        championship_id,
        car_p_id,
        raceclass,
        ROW_NUMBER() OVER (
            PARTITION BY championship_id, raceclass
            ORDER BY points DESC, car_p_id
        ) AS position
    FROM personal_points
),
driver_championship_positions AS (
    SELECT
        cp.championship_id,
        cp.raceclass,
        MIN(cp.position) AS position
    FROM championship_positions cp
    JOIN driver_lineups dl ON dl.car_p_id = cp.car_p_id AND dl.raceclass_name = cp.raceclass
    JOIN driver_championships dc ON dc.championship_id = cp.championship_id
    GROUP BY cp.championship_id, cp.raceclass
)
SELECT
    (SELECT COUNT(DISTINCT race_id) FROM race_results),
    (SELECT COUNT(*) FROM race_results WHERE finish_pos = 1),
    (SELECT COUNT(*) FROM race_results WHERE finish_pos <= 3),
    COALESCE((SELECT SUM(points) FROM race_results), 0),
    (SELECT COUNT(*) FROM qualifying_results WHERE pos = 1),
    (SELECT MIN(finish_pos) FROM race_results),
    (SELECT MIN(pos) FROM qualifying_results),
    (SELECT COUNT(*) FROM driver_championship_positions WHERE position = 1),
    (SELECT MIN(position) FROM driver_championship_positions)
`

const teamStatsQuery = `
WITH team_lineups AS (
    SELECT DISTINCT
        cp.id AS car_p_id,
        rc.name AS raceclass
    FROM car_p cp
    JOIN car c ON c.id = cp.car
    JOIN raceclass rc ON rc.id = c.raceclass
    WHERE cp.team = ?
),
race_results AS (
    SELECT
        r.id AS race_id,
        r.championship AS championship_id,
        cp.id AS car_p_id,
        f.pos AS finish_pos,
        CEIL(
            CASE f.pos
                WHEN 1 THEN 25
                WHEN 2 THEN 18
                WHEN 3 THEN 15
                WHEN 4 THEN 12
                WHEN 5 THEN 10
                WHEN 6 THEN 8
                WHEN 7 THEN 6
                WHEN 8 THEN 4
                WHEN 9 THEN 2
                WHEN 10 THEN 1
                ELSE 0
            END *
            CASE
                WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                WHEN r.type = 2 AND r.duration >= 24 THEN 2
                ELSE 1
            END
        ) AS points
    FROM team_lineups tl
    JOIN car_p cp ON cp.id = tl.car_p_id
    JOIN finish f ON f.car_p = cp.id
    JOIN race r ON r.id = f.race
),
qualifying_results AS (
    SELECT q.pos
    FROM qualifying q
    JOIN team_lineups tl ON tl.car_p_id = q.car_p
),
team_championships AS (
    SELECT DISTINCT championship_id
    FROM race_results
),
personal_points AS (
    SELECT
        r.championship AS championship_id,
        cp.id AS car_p_id,
        cp.team AS team_id,
        rc.name AS raceclass,
        SUM(
            CEIL(
                CASE f.pos
                    WHEN 1 THEN 25
                    WHEN 2 THEN 18
                    WHEN 3 THEN 15
                    WHEN 4 THEN 12
                    WHEN 5 THEN 10
                    WHEN 6 THEN 8
                    WHEN 7 THEN 6
                    WHEN 8 THEN 4
                    WHEN 9 THEN 2
                    WHEN 10 THEN 1
                    ELSE 0
                END *
                CASE
                    WHEN r.type = 2 AND r.duration > 6 AND r.duration < 24 THEN 1.5
                    WHEN r.type = 2 AND r.duration >= 24 THEN 2
                    ELSE 1
                END
            )
        ) AS points
    FROM finish f
    JOIN car_p cp ON f.car_p = cp.id
    JOIN race r ON f.race = r.id
    JOIN car c ON cp.car = c.id
    JOIN raceclass rc ON c.raceclass = rc.id
    GROUP BY r.championship, cp.id, cp.team, rc.name
),
championship_team_points AS (
    SELECT
        championship_id,
        team_id,
        raceclass,
        SUM(points) AS points
    FROM personal_points
    GROUP BY championship_id, team_id, raceclass
),
championship_team_positions AS (
    SELECT
        championship_id,
        team_id,
        raceclass,
        ROW_NUMBER() OVER (
            PARTITION BY championship_id, raceclass
            ORDER BY points DESC, team_id
        ) AS position
    FROM championship_team_points
),
team_positions AS (
    SELECT
        ctp.championship_id,
        ctp.raceclass,
        ctp.position
    FROM championship_team_positions ctp
    JOIN team_championships tc ON tc.championship_id = ctp.championship_id
    WHERE ctp.team_id = ?
)
SELECT
    (SELECT COUNT(DISTINCT race_id) FROM race_results),
    (SELECT COUNT(*) FROM race_results WHERE finish_pos = 1),
    (SELECT COUNT(*) FROM race_results WHERE finish_pos <= 3),
    COALESCE((SELECT SUM(points) FROM race_results), 0),
    (SELECT COUNT(*) FROM qualifying_results WHERE pos = 1),
    (SELECT MIN(finish_pos) FROM race_results),
    (SELECT MIN(pos) FROM qualifying_results),
    (SELECT COUNT(*) FROM team_positions WHERE position = 1),
    (SELECT MIN(position) FROM team_positions)
`
