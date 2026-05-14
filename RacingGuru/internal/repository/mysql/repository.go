package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"RacingGuru/internal/shared"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Repository struct {
	DB *sql.DB
}

type rowScanner interface {
	Scan(dest ...any) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func statementBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Question)
}

func contains(value string) string {
	return "%" + value + "%"
}

func newUUID() uuid.UUID {
	return uuid.New()
}

func dateOnlyExpression(column string) string {
	return fmt.Sprintf("DATE_FORMAT(%s, '%%Y-%%m-%%d')", column)
}

func parseUUID(value string) (uuid.UUID, error) {
	if value == "" {
		return uuid.Nil, nil
	}
	return uuid.Parse(value)
}

func parseNullableUUID(value sql.NullString) (uuid.UUID, error) {
	if !value.Valid {
		return uuid.Nil, nil
	}
	return parseUUID(value.String)
}

func parseDateOnly(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.DateOnly, value)
}

func nullableInt(value sql.NullInt64) int {
	if !value.Valid {
		return 0
	}
	return int(value.Int64)
}

func isNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func ensureRowsAffected(result sql.Result) error {
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return shared.ErrorNotFound
	}
	return nil
}

func (r *Repository) ensureRowsAffectedOrExists(ctx context.Context, result sql.Result, table string, id uuid.UUID) error {
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows > 0 {
		return nil
	}

	exists, err := r.existsByID(ctx, table, id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	return nil
}

func (r *Repository) existsByID(ctx context.Context, table string, id uuid.UUID) (bool, error) {
	query, args, err := statementBuilder().
		Select("1").
		From(table).
		Where(sq.Eq{"id": id.String()}).
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
