package server

import (
	"RacingGuru/internal/core/admin"
	"RacingGuru/internal/core/auth"
	"RacingGuru/internal/core/stats"
	"RacingGuru/internal/core/sudoku"
	"RacingGuru/internal/core/users"
)

type Repo interface {
	admin.Repo
	auth.Repo
	stats.Repo
	sudoku.Repo
	users.Repo
}
