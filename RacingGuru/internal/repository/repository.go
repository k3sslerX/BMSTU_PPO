package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Pool: pool}
}
