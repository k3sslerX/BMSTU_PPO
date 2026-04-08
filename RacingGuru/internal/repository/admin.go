package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO driver (id, name, nationality, birthday) VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id",
		driver.Name, driver.Nationality, driver.Birthday)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return driver, err
	}
	driver.Id = id
	return driver, err
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO team (id, name, country) VALUES (uuid_generate_v4(), $1, $2) RETURNING id",
		team.Name, team.Country)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return team, err
	}
	team.Id = id
	return team, nil
}

func (r *Repository) CreateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO track (id, name, country, lap_length, turns) VALUES (uuid_generate_v4(), $1, $2, $3, $4) RETURNING id",
		track.Name, track.Country, track.Length, track.Turns)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return track, err
	}
	track.Id = id
	return track, nil
}

func (r *Repository) CreateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return carParticipant, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	row := tx.QueryRow(ctx,
		"INSERT INTO car_p (id, car, team, number) VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id",
		carParticipant.CarID, carParticipant.TeamID, carParticipant.Number)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return carParticipant, err
	}
	carParticipant.Id = id

	if err := replaceCarParticipantDrivers(ctx, tx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	if err := tx.Commit(ctx); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO race (id, championship, track, name, date, type, duration) "+
			"VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6) RETURNING id",
		race.ChampionshipId, race.Track.Id, race.Name, race.Date.Format(time.DateOnly), race.Type, race.Duration)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return race, err
	}
	race.Id = id
	return race, nil
}

func (r *Repository) UpsertRaceResult(ctx context.Context, result models.RaceResult) (models.RaceResult, error) {
	if result.RaceID == uuid.Nil || result.CarParticipantID == uuid.Nil ||
		result.FinishPos <= 0 || result.QualifyingPos <= 0 {
		return result, shared.ErrorInvalidData
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return result, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	finishExists, err := upsertRaceStanding(ctx, tx, "finish", result.RaceID, result.CarParticipantID, result.FinishPos)
	if err != nil {
		return result, err
	}
	qualifyingExists, err := upsertRaceStanding(ctx, tx, "qualifying", result.RaceID, result.CarParticipantID, result.QualifyingPos)
	if err != nil {
		return result, err
	}

	if !finishExists && !qualifyingExists {
		row := tx.QueryRow(ctx, "SELECT 1 FROM car_p WHERE id = $1", result.CarParticipantID)
		var exists int
		if err := row.Scan(&exists); err != nil {
			return result, shared.ErrorNotFound
		}

		row = tx.QueryRow(ctx, "SELECT 1 FROM race WHERE id = $1", result.RaceID)
		if err := row.Scan(&exists); err != nil {
			return result, shared.ErrorNotFound
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == uuid.Nil || (driver.Name == "" && driver.Nationality == "" && driver.Birthday == "") {
		return driver, shared.ErrorInvalidData
	}
	sqlString := "UPDATE driver SET "
	fields := make([]any, 0)
	sep := false
	cnt := 1
	if driver.Name != "" {
		sqlString += "name = $" + strconv.Itoa(cnt)
		fields = append(fields, driver.Name)
		sep = true
		cnt++
	}
	if driver.Nationality != "" {
		if sep {
			sqlString += ", "
		}
		sqlString += "nationality = $" + strconv.Itoa(cnt)
		fields = append(fields, driver.Nationality)
		sep = true
		cnt++
	}
	if driver.Birthday != "" {
		if sep {
			sqlString += ", "
		}
		sqlString += "birthday = $" + strconv.Itoa(cnt)
		fields = append(fields, driver.Birthday)
		sep = true
		cnt++
	}
	sqlString += " WHERE id = $" + strconv.Itoa(cnt)
	fields = append(fields, driver.Id)
	res, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return driver, err
	}
	if res.RowsAffected() == 0 {
		return driver, shared.ErrorNotFound
	}
	return driver, nil
}

func (r *Repository) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == uuid.Nil || (team.Name == "" && team.Country == "") {
		return team, shared.ErrorInvalidData
	}
	sqlString := "UPDATE team SET "
	fields := make([]any, 0)
	sep := false
	cnt := 1
	if team.Name != "" {
		sqlString += "name = $" + strconv.Itoa(cnt)
		fields = append(fields, team.Name)
		sep = true
		cnt++
	}
	if team.Country != "" {
		if sep {
			sqlString += ", "
		}
		sqlString += "country = $" + strconv.Itoa(cnt)
		fields = append(fields, team.Country)
		sep = true
		cnt++
	}
	sqlString += " WHERE id = $" + strconv.Itoa(cnt)
	fields = append(fields, team.Id)
	res, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return team, err
	}
	if res.RowsAffected() == 0 {
		return team, shared.ErrorNotFound
	}
	return team, nil
}

func (r *Repository) UpdateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	if track.Id == uuid.Nil || (track.Name == "" && track.Country == "" && track.Length == 0 && track.Turns == 0) {
		return track, shared.ErrorInvalidData
	}

	sqlString := "UPDATE track SET "
	fields := make([]any, 0)
	sep := false
	cnt := 1

	if track.Name != "" {
		sqlString += "name = $" + strconv.Itoa(cnt)
		fields = append(fields, track.Name)
		sep = true
		cnt++
	}
	if track.Country != "" {
		if sep {
			sqlString += ", "
		}
		sqlString += "country = $" + strconv.Itoa(cnt)
		fields = append(fields, track.Country)
		sep = true
		cnt++
	}
	if track.Length != 0 {
		if sep {
			sqlString += ", "
		}
		sqlString += "lap_length = $" + strconv.Itoa(cnt)
		fields = append(fields, track.Length)
		sep = true
		cnt++
	}
	if track.Turns != 0 {
		if sep {
			sqlString += ", "
		}
		sqlString += "turns = $" + strconv.Itoa(cnt)
		fields = append(fields, track.Turns)
		sep = true
		cnt++
	}

	sqlString += " WHERE id = $" + strconv.Itoa(cnt)
	fields = append(fields, track.Id)

	res, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return track, err
	}
	if res.RowsAffected() == 0 {
		return track, shared.ErrorNotFound
	}

	return track, nil
}

func (r *Repository) UpdateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil || (carParticipant.CarID == uuid.Nil && carParticipant.TeamID == uuid.Nil && carParticipant.Number == "") {
		return carParticipant, shared.ErrorInvalidData
	}

	sqlString := "UPDATE car_p SET "
	fields := make([]any, 0)
	sep := false
	cnt := 1

	if carParticipant.CarID != uuid.Nil {
		sqlString += "car = $" + strconv.Itoa(cnt)
		fields = append(fields, carParticipant.CarID)
		sep = true
		cnt++
	}
	if carParticipant.TeamID != uuid.Nil {
		if sep {
			sqlString += ", "
		}
		sqlString += "team = $" + strconv.Itoa(cnt)
		fields = append(fields, carParticipant.TeamID)
		sep = true
		cnt++
	}
	if carParticipant.Number != "" {
		if sep {
			sqlString += ", "
		}
		sqlString += "number = $" + strconv.Itoa(cnt)
		fields = append(fields, carParticipant.Number)
		sep = true
		cnt++
	}

	sqlString += " WHERE id = $" + strconv.Itoa(cnt)
	fields = append(fields, carParticipant.Id)

	res, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return carParticipant, err
	}
	if res.RowsAffected() == 0 {
		return carParticipant, shared.ErrorNotFound
	}

	return carParticipant, nil
}

func (r *Repository) UpdateCarParticipantDrivers(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil {
		return carParticipant, shared.ErrorInvalidData
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return carParticipant, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	row := tx.QueryRow(ctx, "SELECT 1 FROM car_p WHERE id = $1", carParticipant.Id)
	var exists int
	if err := row.Scan(&exists); err != nil {
		return carParticipant, shared.ErrorNotFound
	}

	if err := replaceCarParticipantDrivers(ctx, tx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	if err := tx.Commit(ctx); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == uuid.Nil || (race.Name == "" && race.Date.IsZero() && race.Type == 0 &&
		race.Duration == 0 && race.Track.Id == uuid.Nil && race.ChampionshipId == uuid.Nil) {
		return race, shared.ErrorInvalidData
	}

	sqlString := "UPDATE race SET "
	fields := make([]any, 0)
	sep := false
	cnt := 1

	if race.Name != "" {
		sqlString += "name = $" + strconv.Itoa(cnt)
		fields = append(fields, race.Name)
		sep = true
		cnt++
	}
	if !race.Date.IsZero() {
		if sep {
			sqlString += ", "
		}
		sqlString += "date = $" + strconv.Itoa(cnt)
		fields = append(fields, race.Date)
		sep = true
		cnt++
	}
	if race.Type != 0 {
		if sep {
			sqlString += ", "
		}
		sqlString += "type = $" + strconv.Itoa(cnt)
		fields = append(fields, race.Type)
		sep = true
		cnt++
	}
	if race.Duration != 0 {
		if sep {
			sqlString += ", "
		}
		sqlString += "duration = $" + strconv.Itoa(cnt)
		fields = append(fields, race.Duration)
		sep = true
		cnt++
	}
	if race.Track.Id != uuid.Nil {
		if sep {
			sqlString += ", "
		}
		sqlString += "track = $" + strconv.Itoa(cnt)
		fields = append(fields, race.Track.Id)
		sep = true
		cnt++
	}
	if race.ChampionshipId != uuid.Nil {
		if sep {
			sqlString += ", "
		}
		sqlString += "championship = $" + strconv.Itoa(cnt)
		fields = append(fields, race.ChampionshipId)
		sep = true
		cnt++
	}

	sqlString += " WHERE id = $" + strconv.Itoa(cnt)
	fields = append(fields, race.Id)

	res, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return race, err
	}
	if res.RowsAffected() == 0 {
		return race, shared.ErrorNotFound
	}

	return race, nil
}

func upsertRaceStanding(
	ctx context.Context,
	tx pgx.Tx,
	table string,
	raceID uuid.UUID,
	carParticipantID uuid.UUID,
	position int,
) (bool, error) {
	res, err := tx.Exec(ctx,
		"UPDATE "+table+" SET pos = $1 WHERE race = $2 AND car_p = $3",
		position, raceID, carParticipantID)
	if err != nil {
		return false, err
	}
	if res.RowsAffected() > 0 {
		return true, nil
	}

	_, err = tx.Exec(ctx,
		"INSERT INTO "+table+" (race, car_p, pos) VALUES ($1, $2, $3)",
		raceID, carParticipantID, position)
	if err != nil {
		return false, err
	}

	return false, nil
}

func replaceCarParticipantDrivers(
	ctx context.Context,
	tx pgx.Tx,
	carParticipantID uuid.UUID,
	drivers []models.Driver,
) error {
	if _, err := tx.Exec(ctx, "DELETE FROM team_p WHERE car_p = $1", carParticipantID); err != nil {
		return err
	}

	for _, driver := range drivers {
		if driver.Id == uuid.Nil {
			return shared.ErrorInvalidData
		}
		if _, err := tx.Exec(ctx,
			"INSERT INTO team_p (id, car_p, driver) VALUES (uuid_generate_v4(), $1, $2)",
			carParticipantID, driver.Id); err != nil {
			return err
		}
	}

	return nil
}
