package mysql

import (
	"context"
	"time"

	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *Repository) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	query, args, err := statementBuilder().
		Select("id").
		From("users").
		Where(sq.Eq{"email": user.Email}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var id string
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&id)
	if err == nil {
		return models.User{}, shared.ErrorUserAlreadyExists
	}
	if !isNoRows(err) {
		return models.User{}, err
	}

	user.Id = newUUID()
	query, args, err = statementBuilder().
		Insert("users").
		Columns("id", "name", "email", "passwordHash", "role", "created_at").
		Values(user.Id.String(), user.Name, user.Email, user.Password, user.Role, time.Now().UTC()).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
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
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err == nil {
		return true, nil
	}
	if isNoRows(err) {
		return false, nil
	}

	return false, err
}

func (r *Repository) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	query, args, err := statementBuilder().
		Select("id", "name", "passwordHash", "role").
		From("users").
		Where(sq.Eq{"email": user.Email}).
		ToSql()
	if err != nil {
		return user, err
	}

	var id, name, password, role string
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&id, &name, &password, &role)
	if err != nil {
		if isNoRows(err) {
			return user, shared.ErrorIncorrectPassword
		}
		return user, err
	}
	if password != user.Password {
		return user, shared.ErrorIncorrectPassword
	}

	userID, err := parseUUID(id)
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
		Where(sq.Eq{"id": user.Id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return r.ensureRowsAffectedOrExists(ctx, res, "users", user.Id)
}

func (r *Repository) GetUserByID(ctx context.Context, user models.User) (models.User, error) {
	query, args, err := statementBuilder().
		Select("id", "name", "email", "role").
		From("users").
		Where(sq.Eq{"id": user.Id.String()}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var id, name, email, role string
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&id, &name, &email, &role)
	if err != nil {
		if isNoRows(err) {
			return models.User{}, shared.ErrorNotFound
		}
		return models.User{}, err
	}

	userID, err := parseUUID(id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:    userID,
		Name:  name,
		Email: email,
		Role:  models.Role(role),
	}, nil
}

func (r *Repository) UpdateUserRole(ctx context.Context, user models.User) (models.User, error) {
	if user.Id == uuid.Nil || (user.Role != models.RoleAdmin && user.Role != models.RoleUser) {
		return models.User{}, shared.ErrorInvalidData
	}

	query, args, err := statementBuilder().
		Update("users").
		Set("role", user.Role).
		Where(sq.Eq{"id": user.Id.String()}).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return models.User{}, err
	}
	if err := r.ensureRowsAffectedOrExists(ctx, res, "users", user.Id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) ListFavouriteDrivers(ctx context.Context, user models.User) ([]models.Driver, error) {
	query, args, err := statementBuilder().
		Select("d.id", "d.name", dateOnlyExpression("d.birthday"), "d.nationality").
		From("favourite_drivers fd").
		Join("driver d ON d.id = fd.driver").
		Where(sq.Eq{"fd.user_id": user.Id.String()}).
		OrderBy("d.name ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
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

func (r *Repository) ListFavouriteTeams(ctx context.Context, user models.User) ([]models.Team, error) {
	query, args, err := statementBuilder().
		Select("t.id", "t.name", "t.country").
		From("favourite_teams ft").
		Join("team t ON t.id = ft.team").
		Where(sq.Eq{"ft.user_id": user.Id.String()}).
		OrderBy("t.name ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.QueryContext(ctx, query, args...)
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

func (r *Repository) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	query, args, err := statementBuilder().
		Select("1").
		From("favourite_drivers").
		Where(sq.Eq{"user_id": user.Id.String(), "driver": driver.Id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err == nil {
		query, args, err = statementBuilder().
			Delete("favourite_drivers").
			Where(sq.Eq{"user_id": user.Id.String(), "driver": driver.Id.String()}).
			ToSql()
		if err != nil {
			return err
		}
		_, err = r.DB.ExecContext(ctx, query, args...)
		return err
	}
	if !isNoRows(err) {
		return err
	}

	query, args, err = statementBuilder().
		Insert("favourite_drivers").
		Columns("user_id", "driver").
		Values(user.Id.String(), driver.Id.String()).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r *Repository) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	query, args, err := statementBuilder().
		Select("1").
		From("favourite_teams").
		Where(sq.Eq{"user_id": user.Id.String(), "team": team.Id.String()}).
		ToSql()
	if err != nil {
		return err
	}

	var exists int
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err == nil {
		query, args, err = statementBuilder().
			Delete("favourite_teams").
			Where(sq.Eq{"user_id": user.Id.String(), "team": team.Id.String()}).
			ToSql()
		if err != nil {
			return err
		}
		_, err = r.DB.ExecContext(ctx, query, args...)
		return err
	}
	if !isNoRows(err) {
		return err
	}

	query, args, err = statementBuilder().
		Insert("favourite_teams").
		Columns("user_id", "team").
		Values(user.Id.String(), team.Id.String()).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}
