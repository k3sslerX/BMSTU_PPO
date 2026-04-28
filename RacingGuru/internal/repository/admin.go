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

func (r *Repository) ListCars(ctx context.Context, query string) ([]models.Car, error) {
	builder := statementBuilder().
		Select("id", "model", "year_of_production").
		From("car").
		OrderBy("model ASC", "year_of_production DESC")

	if query != "" {
		builder = builder.Where(sq.ILike{"model": "%" + query + "%"})
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

	cars := make([]models.Car, 0)
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.Id, &car.Model, &car.Year); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *Repository) ListCarParticipants(ctx context.Context, query string) ([]models.CarParticipant, error) {
	builder := statementBuilder().
		Select("id", "car::text", "team::text", "number").
		From("car_p").
		OrderBy("number ASC")

	if query != "" {
		builder = builder.Where(sq.ILike{"number": "%" + query + "%"})
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

	participants := make([]models.CarParticipant, 0)
	for rows.Next() {
		var participant models.CarParticipant
		var carID sql.NullString
		var teamID sql.NullString
		if err := rows.Scan(&participant.Id, &carID, &teamID, &participant.Number); err != nil {
			return nil, err
		}
		if carID.Valid {
			parsedID, err := uuid.Parse(carID.String)
			if err != nil {
				return nil, err
			}
			participant.CarID = parsedID
		}
		if teamID.Valid {
			parsedID, err := uuid.Parse(teamID.String)
			if err != nil {
				return nil, err
			}
			participant.TeamID = parsedID
		}
		participants = append(participants, participant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return participants, nil
}

func (r *Repository) ListChampionships(ctx context.Context, query string) ([]models.Championship, error) {
	builder := statementBuilder().
		Select("id", "year").
		From("championship").
		OrderBy("year DESC")

	if query != "" {
		builder = builder.Where("year::text ILIKE ?", "%"+query+"%")
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

	championships := make([]models.Championship, 0)
	for rows.Next() {
		var championship models.Championship
		if err := rows.Scan(&championship.Id, &championship.Year); err != nil {
			return nil, err
		}
		championships = append(championships, championship)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return championships, nil
}

func (r *Repository) ListRaces(ctx context.Context, query string) ([]models.Race, error) {
	builder := statementBuilder().
		Select("id", "name", "date", "type", "duration", "championship::text", "track::text").
		From("race").
		OrderBy("date DESC", "name ASC")

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

	races := make([]models.Race, 0)
	for rows.Next() {
		var race models.Race
		var championshipID sql.NullString
		var trackID sql.NullString
		if err := rows.Scan(&race.Id, &race.Name, &race.Date, &race.Type, &race.Duration, &championshipID, &trackID); err != nil {
			return nil, err
		}
		if championshipID.Valid {
			parsedID, err := uuid.Parse(championshipID.String)
			if err != nil {
				return nil, err
			}
			race.ChampionshipId = parsedID
		}
		if trackID.Valid {
			parsedID, err := uuid.Parse(trackID.String)
			if err != nil {
				return nil, err
			}
			race.Track.Id = parsedID
		}
		races = append(races, race)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return races, nil
}

func (r *Repository) ListTracks(ctx context.Context, query string) ([]models.Track, error) {
	builder := statementBuilder().
		Select("id", "name", "country", "lap_length", "turns").
		From("track").
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

	tracks := make([]models.Track, 0)
	for rows.Next() {
		var track models.Track
		if err := rows.Scan(&track.Id, &track.Name, &track.Country, &track.Length, &track.Turns); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *Repository) ListUsers(ctx context.Context, query string) ([]models.User, error) {
	builder := statementBuilder().
		Select("id", "name", "email", "role").
		From("users").
		OrderBy("email ASC")

	if query != "" {
		builder = builder.Where(sq.Or{
			sq.ILike{"name": "%" + query + "%"},
			sq.ILike{"email": "%" + query + "%"},
		})
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

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) DeleteDriver(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := deleteByColumn(ctx, tx, "favourite_drivers", "driver", id); err != nil {
		return err
	}
	if err := deleteOneByID(ctx, tx, "driver", id); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}

	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := deleteByColumn(ctx, tx, "favourite_teams", "team", id); err != nil {
		return err
	}
	if err := deleteOneByID(ctx, tx, "team", id); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *Repository) DeleteTrack(ctx context.Context, id uuid.UUID) error {
	return r.deleteOneByID(ctx, "track", id)
}

func (r *Repository) DeleteRace(ctx context.Context, id uuid.UUID) error {
	return r.deleteOneByID(ctx, "race", id)
}

func (r *Repository) DeleteCarParticipant(ctx context.Context, id uuid.UUID) error {
	return r.deleteOneByID(ctx, "car_p", id)
}

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	query, args, err := statementBuilder().
		Insert("driver").
		Columns("id", "name", "nationality", "birthday").
		Values(sq.Expr("uuid_generate_v4()"), driver.Name, driver.Nationality, driver.Birthday).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return driver, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return driver, err
	}
	driver.Id = id
	return driver, err
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	query, args, err := statementBuilder().
		Insert("team").
		Columns("id", "name", "country").
		Values(sq.Expr("uuid_generate_v4()"), team.Name, team.Country).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return team, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return team, err
	}
	team.Id = id
	return team, nil
}

func (r *Repository) CreateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	query, args, err := statementBuilder().
		Insert("track").
		Columns("id", "name", "country", "lap_length", "turns").
		Values(sq.Expr("uuid_generate_v4()"), track.Name, track.Country, track.Length, track.Turns).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return track, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	var id uuid.UUID
	err = row.Scan(&id)
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

	query, args, err := statementBuilder().
		Insert("car_p").
		Columns("id", "car", "team", "number").
		Values(sq.Expr("uuid_generate_v4()"), carParticipant.CarID, carParticipant.TeamID, carParticipant.Number).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return carParticipant, err
	}
	row := tx.QueryRow(ctx, query, args...)
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
	query, args, err := statementBuilder().
		Insert("race").
		Columns("id", "championship", "track", "name", "date", "type", "duration").
		Values(sq.Expr("uuid_generate_v4()"), race.ChampionshipId, race.Track.Id, race.Name, race.Date.Format(time.DateOnly), race.Type, race.Duration).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return race, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	var id uuid.UUID
	err = row.Scan(&id)
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
		query, args, err := statementBuilder().
			Select("1").
			From("car_p").
			Where(sq.Eq{"id": result.CarParticipantID}).
			ToSql()
		if err != nil {
			return result, err
		}
		row := tx.QueryRow(ctx, query, args...)
		var exists int
		if err := row.Scan(&exists); err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return result, err
			}
			return result, shared.ErrorNotFound
		}

		query, args, err = statementBuilder().
			Select("1").
			From("race").
			Where(sq.Eq{"id": result.RaceID}).
			ToSql()
		if err != nil {
			return result, err
		}
		row = tx.QueryRow(ctx, query, args...)
		if err := row.Scan(&exists); err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				return result, err
			}
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
	builder := statementBuilder().Update("driver").Where(sq.Eq{"id": driver.Id})
	if driver.Name != "" {
		builder = builder.Set("name", driver.Name)
	}
	if driver.Nationality != "" {
		builder = builder.Set("nationality", driver.Nationality)
	}
	if driver.Birthday != "" {
		builder = builder.Set("birthday", driver.Birthday)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return driver, err
	}
	res, err := r.Pool.Exec(ctx, query, args...)
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
	builder := statementBuilder().Update("team").Where(sq.Eq{"id": team.Id})
	if team.Name != "" {
		builder = builder.Set("name", team.Name)
	}
	if team.Country != "" {
		builder = builder.Set("country", team.Country)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return team, err
	}
	res, err := r.Pool.Exec(ctx, query, args...)
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

	builder := statementBuilder().Update("track").Where(sq.Eq{"id": track.Id})

	if track.Name != "" {
		builder = builder.Set("name", track.Name)
	}
	if track.Country != "" {
		builder = builder.Set("country", track.Country)
	}
	if track.Length != 0 {
		builder = builder.Set("lap_length", track.Length)
	}
	if track.Turns != 0 {
		builder = builder.Set("turns", track.Turns)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return track, err
	}
	res, err := r.Pool.Exec(ctx, query, args...)
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

	builder := statementBuilder().Update("car_p").Where(sq.Eq{"id": carParticipant.Id})

	if carParticipant.CarID != uuid.Nil {
		builder = builder.Set("car", carParticipant.CarID)
	}
	if carParticipant.TeamID != uuid.Nil {
		builder = builder.Set("team", carParticipant.TeamID)
	}
	if carParticipant.Number != "" {
		builder = builder.Set("number", carParticipant.Number)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return carParticipant, err
	}
	res, err := r.Pool.Exec(ctx, query, args...)
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

	query, args, err := statementBuilder().
		Select("1").
		From("car_p").
		Where(sq.Eq{"id": carParticipant.Id}).
		ToSql()
	if err != nil {
		return carParticipant, err
	}
	row := tx.QueryRow(ctx, query, args...)
	var exists int
	if err := row.Scan(&exists); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return carParticipant, err
		}
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

	builder := statementBuilder().Update("race").Where(sq.Eq{"id": race.Id})

	if race.Name != "" {
		builder = builder.Set("name", race.Name)
	}
	if !race.Date.IsZero() {
		builder = builder.Set("date", race.Date)
	}
	if race.Type != 0 {
		builder = builder.Set("type", race.Type)
	}
	if race.Duration != 0 {
		builder = builder.Set("duration", race.Duration)
	}
	if race.Track.Id != uuid.Nil {
		builder = builder.Set("track", race.Track.Id)
	}
	if race.ChampionshipId != uuid.Nil {
		builder = builder.Set("championship", race.ChampionshipId)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return race, err
	}
	res, err := r.Pool.Exec(ctx, query, args...)
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
	query, args, err := statementBuilder().
		Update(table).
		Set("pos", position).
		Where(sq.Eq{"race": raceID, "car_p": carParticipantID}).
		ToSql()
	if err != nil {
		return false, err
	}
	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return false, err
	}
	if res.RowsAffected() > 0 {
		return true, nil
	}

	query, args, err = statementBuilder().
		Insert(table).
		Columns("race", "car_p", "pos").
		Values(raceID, carParticipantID, position).
		ToSql()
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(ctx, query, args...)
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
	query, args, err := statementBuilder().
		Delete("team_p").
		Where(sq.Eq{"car_p": carParticipantID}).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return err
	}

	for _, driver := range drivers {
		if driver.Id == uuid.Nil {
			return shared.ErrorInvalidData
		}
		query, args, err = statementBuilder().
			Insert("team_p").
			Columns("car_p", "driver").
			Values(carParticipantID, driver.Id).
			ToSql()
		if err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, query, args...); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) deleteOneByID(ctx context.Context, table string, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}

	query, args, err := statementBuilder().
		Delete(table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return shared.ErrorNotFound
	}

	return nil
}

func deleteOneByID(ctx context.Context, tx pgx.Tx, table string, id uuid.UUID) error {
	query, args, err := statementBuilder().
		Delete(table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return shared.ErrorNotFound
	}

	return nil
}

func deleteByColumn(ctx context.Context, tx pgx.Tx, table string, column string, id uuid.UUID) error {
	query, args, err := statementBuilder().
		Delete(table).
		Where(sq.Eq{column: id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	return err
}
