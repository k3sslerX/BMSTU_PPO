package repository

import (
	"RacingGuru/internal/models"
	"context"
	"fmt"

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

func (r *Repository) getDriversByConditions(ctx context.Context, firstCondition models.SudokuCondition, secondCondition models.SudokuCondition) ([]models.Driver, error) {
	firstConditionSQL, firstArg, err := buildStatsConditionSQL("stats", firstCondition)
	if err != nil {
		return nil, err
	}
	secondConditionSQL, secondArg, err := buildStatsConditionSQL("stats", secondCondition)
	if err != nil {
		return nil, err
	}

	query, args, err := statementBuilder().
		Select("d.id", "d.name", "d.birthday::text", "d.nationality").
		Distinct().
		From("driver d").
		JoinClause("CROSS JOIN LATERAL CalculateDriverStats(d.id) stats").
		Where(sq.Expr(firstConditionSQL, firstArg)).
		Where(sq.Expr(secondConditionSQL, secondArg)).
		OrderBy("d.name").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	drivers := make([]models.Driver, 0)
	for rows.Next() {
		driver := models.Driver{}
		if err := rows.Scan(&driver.Id, &driver.Name, &driver.Birthday, &driver.Nationality); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

func (r *Repository) getTeamsByConditions(ctx context.Context, firstCondition models.SudokuCondition, secondCondition models.SudokuCondition) ([]models.Team, error) {
	firstConditionSQL, firstArg, err := buildStatsConditionSQL("stats", firstCondition)
	if err != nil {
		return nil, err
	}
	secondConditionSQL, secondArg, err := buildStatsConditionSQL("stats", secondCondition)
	if err != nil {
		return nil, err
	}

	query, args, err := statementBuilder().
		Select("t.id", "t.name", "t.country").
		Distinct().
		From("team t").
		JoinClause("CROSS JOIN LATERAL CalculateTeamStats(t.id) stats").
		Where(sq.Expr(firstConditionSQL, firstArg)).
		Where(sq.Expr(secondConditionSQL, secondArg)).
		OrderBy("t.name").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make([]models.Team, 0)
	for rows.Next() {
		team := models.Team{}
		if err := rows.Scan(&team.Id, &team.Name, &team.Country); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func buildStatsConditionSQL(statsAlias string, condition models.SudokuCondition) (string, any, error) {
	fieldName, err := mapStatsField(condition.Field)
	if err != nil {
		return "", nil, err
	}

	operator, err := mapConditionOperator(condition.Op)
	if err != nil {
		return "", nil, err
	}

	return fmt.Sprintf("%s.%s %s ?", statsAlias, fieldName, operator), condition.Value, nil
}

func mapStatsField(field string) (string, error) {
	switch field {
	case "total_races":
		return "total_races", nil
	case "total_wins":
		return "total_wins", nil
	case "total_podiums":
		return "total_podiums", nil
	case "total_points":
		return "total_points", nil
	case "total_poles":
		return "total_poles", nil
	case "best_finish":
		return "best_finish", nil
	case "best_qualifying":
		return "best_qualifying", nil
	case "championships_wins":
		return "championship_wins", nil
	case "best_championship_position":
		return "best_championship_position", nil
	default:
		return "", fmt.Errorf("unsupported sudoku stats field: %s", field)
	}
}

func mapConditionOperator(operator models.ConditionOperator) (string, error) {
	switch operator {
	case models.ConditionOperatorGT:
		return ">", nil
	case models.ConditionOperatorLT:
		return "<", nil
	case models.ConditionOperatorGTE:
		return ">=", nil
	case models.ConditionOperatorLTE:
		return "<=", nil
	case models.ConditionOperatorEQ:
		return "=", nil
	default:
		return "", fmt.Errorf("unsupported sudoku condition operator: %s", operator)
	}
}
