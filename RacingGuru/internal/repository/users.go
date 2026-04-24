package repository

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (r *Repository) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	var id string
	query, args, err := statementBuilder().
		Select("id").
		From("users").
		Where(sq.Eq{"email": user.Email}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&id)
	if err == nil {
		return models.User{}, shared.ErrorUserAlreadyExists
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return models.User{}, err
	}
	query, args, err = statementBuilder().
		Insert("users").
		Columns("id", "name", "email", "passwordHash", "role", "created_at").
		Values(sq.Expr("uuid_generate_v4()"), user.Name, user.Email, user.Password, user.Role, time.Now().UTC()).
		ToSql()
	if err != nil {
		return models.User{}, err
	}
	_, err = r.Pool.Exec(ctx, query, args...)
	return user, err
}

func (r *Repository) HasAdmin(ctx context.Context) (bool, error) {
	query, args, err := statementBuilder().
		Select("1").
		From("users").
		Where(sq.Eq{"role": models.RoleAdmin}).
		Limit(1).
		ToSql()
	if err != nil {
		return false, err
	}

	var exists int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&exists)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

	return false, err
}

func (r *Repository) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	var id, name, password, role string
	query, args, err := statementBuilder().
		Select("id", "name", "passwordHash", "role").
		From("users").
		Where(sq.Eq{"email": user.Email}).
		ToSql()
	if err != nil {
		return user, err
	}
	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&id, &name, &password, &role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, shared.ErrorIncorrectPassword
		}
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

func (r *Repository) UserChangePassword(ctx context.Context, user models.User, pwd string) error {
	query, args, err := statementBuilder().
		Update("users").
		Set("passwordHash", pwd).
		Where(sq.Eq{"id": user.Id}).
		ToSql()
	if err != nil {
		return err
	}

	tag, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return shared.ErrorNotFound
	}

	return nil
}

func (r *Repository) UpdateUserRole(ctx context.Context, user models.User) (models.User, error) {
	if user.Id == uuid.Nil || (user.Role != models.RoleAdmin && user.Role != models.RoleUser) {
		return models.User{}, shared.ErrorInvalidData
	}

	query, args, err := statementBuilder().
		Update("users").
		Set("role", user.Role).
		Where(sq.Eq{"id": user.Id}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	tag, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return models.User{}, err
	}
	if tag.RowsAffected() == 0 {
		return models.User{}, shared.ErrorNotFound
	}

	return user, nil
}

func (r *Repository) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	query, args, err := statementBuilder().
		Select("1").
		From("favourite_drivers").
		Where(sq.Eq{"user_id": user.Id, "driver": driver.Id}).
		ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&exists)
	if err == nil {
		query, args, err = statementBuilder().
			Delete("favourite_drivers").
			Where(sq.Eq{"user_id": user.Id, "driver": driver.Id}).
			ToSql()
		if err != nil {
			return err
		}
		_, err = r.Pool.Exec(ctx, query, args...)
		return err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	query, args, err = statementBuilder().
		Insert("favourite_drivers").
		Columns("user_id", "driver").
		Values(user.Id, driver.Id).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, query, args...)
	return err
}

func (r *Repository) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	query, args, err := statementBuilder().
		Select("1").
		From("favourite_teams").
		Where(sq.Eq{"user_id": user.Id, "team": team.Id}).
		ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = r.Pool.QueryRow(ctx, query, args...).Scan(&exists)
	if err == nil {
		query, args, err = statementBuilder().
			Delete("favourite_teams").
			Where(sq.Eq{"user_id": user.Id, "team": team.Id}).
			ToSql()
		if err != nil {
			return err
		}
		_, err = r.Pool.Exec(ctx, query, args...)
		return err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	query, args, err = statementBuilder().
		Insert("favourite_teams").
		Columns("user_id", "team").
		Values(user.Id, team.Id).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.Pool.Exec(ctx, query, args...)
	return err
}
