package mysql

import (
	"context"
	"database/sql"

	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *Repository) ListCars(ctx context.Context, query string) ([]models.Car, error) {
	builder := statementBuilder().
		Select("c.id", "c.model", "c.year_of_production", "COALESCE(rc.name, '')").
		From("car c").
		LeftJoin("raceclass rc ON rc.id = c.raceclass").
		OrderBy("c.model ASC", "c.year_of_production DESC")

	if query != "" {
		builder = builder.Where(sq.Or{
			sq.Like{"c.model": contains(query)},
			sq.Like{"rc.name": contains(query)},
			sq.Expr("CAST(c.year_of_production AS CHAR) LIKE ?", contains(query)),
		})
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

	cars := make([]models.Car, 0)
	for rows.Next() {
		var car models.Car
		var id string
		if err := rows.Scan(&id, &car.Model, &car.Year, &car.RaceClass); err != nil {
			return nil, err
		}
		car.Id, err = parseUUID(id)
		if err != nil {
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
		Select("cp.id", "cp.car", "cp.team", "cp.number", "d.id", "d.name").
		From("car_p cp").
		LeftJoin("team_p tp ON tp.car_p = cp.id").
		LeftJoin("driver d ON d.id = tp.driver").
		OrderBy("cp.number ASC", "d.name ASC")

	if query != "" {
		builder = builder.Where(sq.Like{"cp.number": contains(query)})
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

	participantsByID := make(map[uuid.UUID]int)
	participants := make([]models.CarParticipant, 0)
	for rows.Next() {
		var id string
		var carID, teamID, driverID, driverName sql.NullString
		var number string
		if err := rows.Scan(&id, &carID, &teamID, &number, &driverID, &driverName); err != nil {
			return nil, err
		}

		participantID, err := parseUUID(id)
		if err != nil {
			return nil, err
		}

		participantIdx, ok := participantsByID[participantID]
		if !ok {
			parsedCarID, err := parseNullableUUID(carID)
			if err != nil {
				return nil, err
			}
			parsedTeamID, err := parseNullableUUID(teamID)
			if err != nil {
				return nil, err
			}

			participants = append(participants, models.CarParticipant{
				Id:      participantID,
				CarID:   parsedCarID,
				TeamID:  parsedTeamID,
				Number:  number,
				Drivers: make([]models.Driver, 0),
			})
			participantIdx = len(participants) - 1
			participantsByID[participantID] = participantIdx
		}

		if driverID.Valid {
			parsedDriverID, err := parseUUID(driverID.String)
			if err != nil {
				return nil, err
			}
			participants[participantIdx].Drivers = append(participants[participantIdx].Drivers, models.Driver{
				Id:   parsedDriverID,
				Name: driverName.String,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return participants, nil
}

func (r *Repository) ListChampionships(ctx context.Context, query string) ([]models.Championship, error) {
	builder := statementBuilder().
		Select("c.id", "c.year", "COALESCE(o.name, '')").
		From("championship c").
		LeftJoin("organizer o ON o.id = c.organizer").
		OrderBy("c.year DESC")

	if query != "" {
		builder = builder.Where(sq.Or{
			sq.Expr("CAST(c.year AS CHAR) LIKE ?", contains(query)),
			sq.Like{"o.name": contains(query)},
		})
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

	championships := make([]models.Championship, 0)
	for rows.Next() {
		var championship models.Championship
		var id string
		if err := rows.Scan(&id, &championship.Year, &championship.Organizer); err != nil {
			return nil, err
		}
		championship.Id, err = parseUUID(id)
		if err != nil {
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
		Select("id", "name", dateOnlyExpression("date"), "type", "duration", "championship", "track").
		From("race").
		OrderBy("date DESC", "name ASC")

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

	races := make([]models.Race, 0)
	for rows.Next() {
		var race models.Race
		var id, date string
		var championshipID, trackID sql.NullString
		if err := rows.Scan(&id, &race.Name, &date, &race.Type, &race.Duration, &championshipID, &trackID); err != nil {
			return nil, err
		}
		race.Id, err = parseUUID(id)
		if err != nil {
			return nil, err
		}
		race.Date, err = parseDateOnly(date)
		if err != nil {
			return nil, err
		}
		race.ChampionshipId, err = parseNullableUUID(championshipID)
		if err != nil {
			return nil, err
		}
		race.Track.Id, err = parseNullableUUID(trackID)
		if err != nil {
			return nil, err
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

	tracks := make([]models.Track, 0)
	for rows.Next() {
		var track models.Track
		var id string
		if err := rows.Scan(&id, &track.Name, &track.Country, &track.Length, &track.Turns); err != nil {
			return nil, err
		}
		track.Id, err = parseUUID(id)
		if err != nil {
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
			sq.Like{"name": contains(query)},
			sq.Like{"email": contains(query)},
		})
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

	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		var id, role string
		if err := rows.Scan(&id, &user.Name, &user.Email, &role); err != nil {
			return nil, err
		}
		user.Id, err = parseUUID(id)
		if err != nil {
			return nil, err
		}
		user.Role = models.Role(role)
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

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := deleteByColumn(ctx, tx, "favourite_drivers", "driver", id); err != nil {
		return err
	}
	if err := deleteOneByID(ctx, tx, "driver", id); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repository) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := deleteByColumn(ctx, tx, "favourite_teams", "team", id); err != nil {
		return err
	}
	if err := deleteOneByID(ctx, tx, "team", id); err != nil {
		return err
	}

	return tx.Commit()
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
	driver.Id = newUUID()
	query, args, err := statementBuilder().
		Insert("driver").
		Columns("id", "name", "nationality", "birthday").
		Values(driver.Id.String(), driver.Name, driver.Nationality, driver.Birthday).
		ToSql()
	if err != nil {
		return driver, err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return driver, err
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	team.Id = newUUID()
	query, args, err := statementBuilder().
		Insert("team").
		Columns("id", "name", "country").
		Values(team.Id.String(), team.Name, team.Country).
		ToSql()
	if err != nil {
		return team, err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return team, err
}

func (r *Repository) CreateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	track.Id = newUUID()
	query, args, err := statementBuilder().
		Insert("track").
		Columns("id", "name", "country", "lap_length", "turns").
		Values(track.Id.String(), track.Name, track.Country, track.Length, track.Turns).
		ToSql()
	if err != nil {
		return track, err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return track, err
}

func (r *Repository) CreateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return carParticipant, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	carParticipant.Id = newUUID()
	query, args, err := statementBuilder().
		Insert("car_p").
		Columns("id", "car", "team", "number").
		Values(carParticipant.Id.String(), carParticipant.CarID.String(), carParticipant.TeamID.String(), carParticipant.Number).
		ToSql()
	if err != nil {
		return carParticipant, err
	}
	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return carParticipant, err
	}

	if err := replaceCarParticipantDrivers(ctx, tx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	if err := tx.Commit(); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	race.Id = newUUID()
	query, args, err := statementBuilder().
		Insert("race").
		Columns("id", "championship", "track", "name", "date", "type", "duration").
		Values(race.Id.String(), race.ChampionshipId.String(), race.Track.Id.String(), race.Name, race.Date.Format("2006-01-02"), race.Type, race.Duration).
		ToSql()
	if err != nil {
		return race, err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return race, err
}

func (r *Repository) UpsertRaceResult(ctx context.Context, result models.RaceResult) (models.RaceResult, error) {
	if result.RaceID == uuid.Nil || result.CarParticipantID == uuid.Nil ||
		result.FinishPos <= 0 || result.QualifyingPos <= 0 {
		return result, shared.ErrorInvalidData
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return result, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := ensureTxExistsByID(ctx, tx, "car_p", result.CarParticipantID); err != nil {
		return result, err
	}
	if err := ensureTxExistsByID(ctx, tx, "race", result.RaceID); err != nil {
		return result, err
	}

	if _, err := upsertRaceStanding(ctx, tx, "finish", result.RaceID, result.CarParticipantID, result.FinishPos); err != nil {
		return result, err
	}
	if _, err := upsertRaceStanding(ctx, tx, "qualifying", result.RaceID, result.CarParticipantID, result.QualifyingPos); err != nil {
		return result, err
	}

	if err := tx.Commit(); err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == uuid.Nil || (driver.Name == "" && driver.Nationality == "" && driver.Birthday == "") {
		return driver, shared.ErrorInvalidData
	}

	builder := statementBuilder().Update("driver").Where(sq.Eq{"id": driver.Id.String()})
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
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return driver, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "driver", driver.Id); err != nil {
		return driver, err
	}

	return driver, nil
}

func (r *Repository) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == uuid.Nil || (team.Name == "" && team.Country == "") {
		return team, shared.ErrorInvalidData
	}

	builder := statementBuilder().Update("team").Where(sq.Eq{"id": team.Id.String()})
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
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return team, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "team", team.Id); err != nil {
		return team, err
	}

	return team, nil
}

func (r *Repository) UpdateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	if track.Id == uuid.Nil || (track.Name == "" && track.Country == "" && track.Length == 0 && track.Turns == 0) {
		return track, shared.ErrorInvalidData
	}

	builder := statementBuilder().Update("track").Where(sq.Eq{"id": track.Id.String()})
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
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return track, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "track", track.Id); err != nil {
		return track, err
	}

	return track, nil
}

func (r *Repository) UpdateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil || (carParticipant.CarID == uuid.Nil && carParticipant.TeamID == uuid.Nil && carParticipant.Number == "") {
		return carParticipant, shared.ErrorInvalidData
	}

	builder := statementBuilder().Update("car_p").Where(sq.Eq{"id": carParticipant.Id.String()})
	if carParticipant.CarID != uuid.Nil {
		builder = builder.Set("car", carParticipant.CarID.String())
	}
	if carParticipant.TeamID != uuid.Nil {
		builder = builder.Set("team", carParticipant.TeamID.String())
	}
	if carParticipant.Number != "" {
		builder = builder.Set("number", carParticipant.Number)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return carParticipant, err
	}
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return carParticipant, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "car_p", carParticipant.Id); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) UpdateCarParticipantDrivers(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil {
		return carParticipant, shared.ErrorInvalidData
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return carParticipant, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := ensureTxExistsByID(ctx, tx, "car_p", carParticipant.Id); err != nil {
		return carParticipant, err
	}

	if err := replaceCarParticipantDrivers(ctx, tx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	if err := tx.Commit(); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == uuid.Nil || (race.Name == "" && race.Date.IsZero() && race.Type == 0 &&
		race.Duration == 0 && race.Track.Id == uuid.Nil && race.ChampionshipId == uuid.Nil) {
		return race, shared.ErrorInvalidData
	}

	builder := statementBuilder().Update("race").Where(sq.Eq{"id": race.Id.String()})
	if race.Name != "" {
		builder = builder.Set("name", race.Name)
	}
	if !race.Date.IsZero() {
		builder = builder.Set("date", race.Date.Format("2006-01-02"))
	}
	if race.Type != 0 {
		builder = builder.Set("type", race.Type)
	}
	if race.Duration != 0 {
		builder = builder.Set("duration", race.Duration)
	}
	if race.Track.Id != uuid.Nil {
		builder = builder.Set("track", race.Track.Id.String())
	}
	if race.ChampionshipId != uuid.Nil {
		builder = builder.Set("championship", race.ChampionshipId.String())
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return race, err
	}
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return race, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "race", race.Id); err != nil {
		return race, err
	}

	return race, nil
}

func upsertRaceStanding(
	ctx context.Context,
	tx *sql.Tx,
	table string,
	raceID uuid.UUID,
	carParticipantID uuid.UUID,
	position int,
) (bool, error) {
	query, args, err := statementBuilder().
		Select("1").
		From(table).
		Where(sq.Eq{"race": raceID.String(), "car_p": carParticipantID.String()}).
		Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	var exists int
	err = tx.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil && !isNoRows(err) {
		return false, err
	}
	if err == nil {
		query, args, err = statementBuilder().
			Update(table).
			Set("pos", position).
			Where(sq.Eq{"race": raceID.String(), "car_p": carParticipantID.String()}).
			ToSql()
		if err != nil {
			return false, err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return true, err
	}

	query, args, err = statementBuilder().
		Insert(table).
		Columns("race", "car_p", "pos").
		Values(raceID.String(), carParticipantID.String(), position).
		ToSql()
	if err != nil {
		return false, err
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return false, err
	}

	return false, nil
}

func replaceCarParticipantDrivers(
	ctx context.Context,
	tx *sql.Tx,
	carParticipantID uuid.UUID,
	drivers []models.Driver,
) error {
	query, args, err := statementBuilder().
		Delete("team_p").
		Where(sq.Eq{"car_p": carParticipantID.String()}).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	for _, driver := range drivers {
		if driver.Id == uuid.Nil {
			return shared.ErrorInvalidData
		}
		query, args, err = statementBuilder().
			Insert("team_p").
			Columns("car_p", "driver").
			Values(carParticipantID.String(), driver.Id.String()).
			ToSql()
		if err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
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
		Where(sq.Eq{"id": id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return ensureRowsAffected(res)
}

func deleteOneByID(ctx context.Context, tx *sql.Tx, table string, id uuid.UUID) error {
	query, args, err := statementBuilder().
		Delete(table).
		Where(sq.Eq{"id": id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return ensureRowsAffected(res)
}

func deleteByColumn(ctx context.Context, tx *sql.Tx, table string, column string, id uuid.UUID) error {
	query, args, err := statementBuilder().
		Delete(table).
		Where(sq.Eq{column: id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	return err
}

func ensureTxExistsByID(ctx context.Context, tx *sql.Tx, table string, id uuid.UUID) error {
	query, args, err := statementBuilder().
		Select("1").
		From(table).
		Where(sq.Eq{"id": id.String()}).
		Limit(1).
		ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = tx.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err == nil {
		return nil
	}
	if isNoRows(err) {
		return shared.ErrorNotFound
	}
	return err
}
