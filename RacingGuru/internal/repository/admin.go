package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"strconv"

	"github.com/google/uuid"
)

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO driver (name, nationality, birthday) VALUES ($1, $2, $3) RETURNING id",
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
		"INSERT INTO team (name, country) VALUES ($1, $2) RETURNING id",
		team.Name, team.Country)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return team, err
	}
	team.Id = id
	return team, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
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
	_, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return driver, err
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
	_, err := r.Pool.Exec(ctx, sqlString, fields...)
	if err != nil {
		return team, err
	}
	return team, nil
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
}
