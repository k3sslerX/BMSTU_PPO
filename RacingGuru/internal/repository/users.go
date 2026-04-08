package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	var id string
	row := r.Pool.QueryRow(ctx,
		"SELECT id FROM users WHERE email = $1", user.Email)
	err := row.Scan(&id)
	if err == nil {
		return models.User{}, shared.ErrorUserAlreadyExists
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return models.User{}, err
	}
	_, err = r.Pool.Exec(ctx,
		"INSERT INTO users (id, name, email, passwordHash, role, created_at) VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)",
		user.Name, user.Email, user.Password, user.Role, time.Now().UTC())
	return user, err
}

func (r *Repository) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	var id, name, password, role string
	row := r.Pool.QueryRow(ctx,
		"SELECT id, name, password, role FROM users WHERE email = $1", user.Email)
	err := row.Scan(&id, &name, &password, &role)
	if err != nil {
		return user, err
	}
	if password != user.Password {
		return user, shared.ErrorIncorrectPassword
	}
	userID, err := uuid.Parse(id)
	if err != nil {
		return user, err
	}
	user.Id = userID
	user.Name = name
	user.Role = models.Role(role)

	return user, nil
}

func (r *Repository) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	var exists bool
	err := r.Pool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM favourite_drivers WHERE user_id = $1 AND driver = $2)",
		user.Id, driver.Id,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = r.Pool.Exec(ctx,
			"DELETE FROM favourite_drivers WHERE user_id = $1 AND driver = $2",
			user.Id, driver.Id,
		)
		return err
	}

	_, err = r.Pool.Exec(ctx,
		"INSERT INTO favourite_drivers (user_id, driver) VALUES ($1, $2)",
		user.Id, driver.Id,
	)
	return err
}

func (r *Repository) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	var exists bool
	err := r.Pool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM favourite_teams WHERE user_id = $1 AND team = $2)",
		user.Id, team.Id,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		_, err = r.Pool.Exec(ctx,
			"DELETE FROM favourite_teams WHERE user_id = $1 AND team = $2",
			user.Id, team.Id,
		)
		return err
	}

	_, err = r.Pool.Exec(ctx,
		"INSERT INTO favourite_teams (user_id, team) VALUES ($1, $2)",
		user.Id, team.Id,
	)
	return err
}
