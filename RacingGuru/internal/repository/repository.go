package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func statementBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Pool: pool}
}
