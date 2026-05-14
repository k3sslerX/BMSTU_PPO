package mysql

import (
	"context"
	"fmt"

	"RacingGuru/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) GetDriverMatrix(ctx context.Context, matrix models.MatrixDrivers) (models.MatrixDrivers, error) {
	for colIdx, colCondition := range matrix.ConditionSpecs[0] {
		for rowIdx, rowCondition := range matrix.ConditionSpecs[1] {
			drivers, err := r.getDriversByConditions(ctx, rowCondition, colCondition)
			if err != nil {
				return models.MatrixDrivers{}, err
			}
			matrix.Field[rowIdx][colIdx] = drivers
		}
	}

	return matrix, nil
}

func (r *Repository) GetTeamMatrix(ctx context.Context, matrix models.MatrixTeams) (models.MatrixTeams, error) {
	for colIdx, colCondition := range matrix.ConditionSpecs[0] {
		for rowIdx, rowCondition := range matrix.ConditionSpecs[1] {
			teams, err := r.getTeamsByConditions(ctx, rowCondition, colCondition)
			if err != nil {
				return models.MatrixTeams{}, err
			}
			matrix.Field[rowIdx][colIdx] = teams
		}
	}

	return matrix, nil
}

func (r *Repository) CompleteSudokuMatrix(ctx context.Context, user models.User, matrixType models.SudokuMatrixType) error {
	query, args, err := statementBuilder().
		Insert("sudoku_matrix_completions").
		Columns("user_id", "matrix_type", "completed_on").
		Values(user.Id.String(), string(matrixType), sq.Expr("CURRENT_DATE")).
		Suffix("ON DUPLICATE KEY UPDATE completed_on = completed_on").
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r *Repository) GetSudokuCompletionStats(ctx context.Context, user models.User) (models.SudokuCompletionStats, error) {
	query, args, err := statementBuilder().
		Select(
			"COALESCE(SUM(CASE WHEN matrix_type = 'drivers' THEN 1 ELSE 0 END), 0)",
			"COALESCE(SUM(CASE WHEN matrix_type = 'teams' THEN 1 ELSE 0 END), 0)",
		).
		From("sudoku_matrix_completions").
		Where(sq.Eq{"user_id": user.Id.String()}).
		ToSql()
	if err != nil {
		return models.SudokuCompletionStats{}, err
	}

	var stats models.SudokuCompletionStats
	err = r.DB.QueryRowContext(ctx, query, args...).Scan(&stats.DriverMatrices, &stats.TeamMatrices)
	if err != nil {
		return models.SudokuCompletionStats{}, err
	}

	return stats, nil
}

func (r *Repository) getDriversByConditions(ctx context.Context, firstCondition models.SudokuCondition, secondCondition models.SudokuCondition) ([]models.Driver, error) {
	drivers, err := r.ListDrivers(ctx, "")
	if err != nil {
		return nil, err
	}

	matched := make([]models.Driver, 0)
	for _, driver := range drivers {
		driverStats, err := r.GetDriverStats(ctx, driver)
		if err != nil {
			return nil, err
		}
		if matchStatsCondition(driverStats.Stats, firstCondition) && matchStatsCondition(driverStats.Stats, secondCondition) {
			matched = append(matched, driver)
		}
	}

	return matched, nil
}

func (r *Repository) getTeamsByConditions(ctx context.Context, firstCondition models.SudokuCondition, secondCondition models.SudokuCondition) ([]models.Team, error) {
	teams, err := r.ListTeams(ctx, "")
	if err != nil {
		return nil, err
	}

	matched := make([]models.Team, 0)
	for _, team := range teams {
		teamStats, err := r.GetTeamStats(ctx, team)
		if err != nil {
			return nil, err
		}
		if matchStatsCondition(teamStats.Stats, firstCondition) && matchStatsCondition(teamStats.Stats, secondCondition) {
			matched = append(matched, team)
		}
	}

	return matched, nil
}

func matchStatsCondition(stats models.Stats, condition models.SudokuCondition) bool {
	value, err := statsValue(stats, condition.Field)
	if err != nil {
		return false
	}

	switch condition.Op {
	case models.ConditionOperatorGT:
		return value > condition.Value
	case models.ConditionOperatorLT:
		return value < condition.Value
	case models.ConditionOperatorGTE:
		return value >= condition.Value
	case models.ConditionOperatorLTE:
		return value <= condition.Value
	case models.ConditionOperatorEQ:
		return value == condition.Value
	default:
		return false
	}
}

func statsValue(stats models.Stats, field string) (int, error) {
	switch field {
	case "total_races":
		return stats.TotalRaces, nil
	case "total_wins":
		return stats.TotalWins, nil
	case "total_podiums":
		return stats.TotalPodiums, nil
	case "total_points":
		return stats.TotalPoints, nil
	case "total_poles":
		return stats.TotalPoles, nil
	case "best_finish":
		return stats.BestFinish, nil
	case "best_qualifying":
		return stats.BestQualifying, nil
	case "championships_wins":
		return stats.ChampionshipsWins, nil
	case "best_championship_position":
		return stats.BestChampionshipPosition, nil
	default:
		return 0, fmt.Errorf("unsupported sudoku stats field: %s", field)
	}
}
