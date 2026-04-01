package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"strconv"
)

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO driver (id, name, nationality, birthday) VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id",
		driver.Name, driver.Nationality, driver.Birthday)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return driver, err
	}
	driver.Id = models.Uuid(id)
	return driver, err
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	row := r.Pool.QueryRow(ctx,
		"INSERT INTO team (id, name, country) VALUES (uuid_generate_v4(), $1, $2) RETURNING id",
		team.Name, team.Country)
	var id string
	err := row.Scan(&id)
	if err != nil {
		return team, err
	}
	team.Id = models.Uuid(id)
	return team, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
}

func (r *Repository) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == "" || (driver.Name == "" && driver.Nationality == "" && driver.Birthday == "") {
		return driver, shared.ErrorInvalidData
	}
	sqlString := "UPDATE driver SET "
	fields := make([]string, 0)
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
	fields = append(fields, string(driver.Id))
	_, err := r.Pool.Exec(ctx, sqlString, fields)
	if err != nil {
		return driver, err
	}
	return driver, nil
}

func (r *Repository) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == "" || (team.Name == "" && team.Country == "") {
		return team, shared.ErrorInvalidData
	}
	sqlString := "UPDATE team SET "
	fields := make([]string, 0)
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
	fields = append(fields, string(team.Id))
	_, err := r.Pool.Exec(ctx, sqlString, fields)
	if err != nil {
		return team, err
	}
	return team, nil
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	return race, nil
}
