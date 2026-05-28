package mongodb

import (
	"RacingGuru/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetDriverMatrix(ctx context.Context, matrix models.MatrixDrivers) (models.MatrixDrivers, error) {
	data, err := r.loadStatsData(ctx)
	if err != nil {
		return models.MatrixDrivers{}, err
	}

	drivers := make([]driverDoc, 0, len(data.drivers))
	for _, driver := range data.drivers {
		drivers = append(drivers, driver)
	}

	for colIdx, colCondition := range matrix.ConditionSpecs[0] {
		for rowIdx, rowCondition := range matrix.ConditionSpecs[1] {
			cell := make([]models.Driver, 0)
			for _, driver := range drivers {
				stats := data.driverStats(driver.ID)
				ok, err := matchesConditions(stats, rowCondition, colCondition)
				if err != nil {
					return models.MatrixDrivers{}, err
				}
				if !ok {
					continue
				}
				model, err := driver.model()
				if err != nil {
					return models.MatrixDrivers{}, err
				}
				cell = append(cell, model)
			}
			matrix.Field[rowIdx][colIdx] = cell
		}
	}

	return matrix, nil
}

func (r *Repository) GetTeamMatrix(ctx context.Context, matrix models.MatrixTeams) (models.MatrixTeams, error) {
	data, err := r.loadStatsData(ctx)
	if err != nil {
		return models.MatrixTeams{}, err
	}

	teams := make([]teamDoc, 0, len(data.teams))
	for _, team := range data.teams {
		teams = append(teams, team)
	}

	for colIdx, colCondition := range matrix.ConditionSpecs[0] {
		for rowIdx, rowCondition := range matrix.ConditionSpecs[1] {
			cell := make([]models.Team, 0)
			for _, team := range teams {
				stats := data.teamStats(team.ID)
				ok, err := matchesConditions(stats, rowCondition, colCondition)
				if err != nil {
					return models.MatrixTeams{}, err
				}
				if !ok {
					continue
				}
				model, err := team.model()
				if err != nil {
					return models.MatrixTeams{}, err
				}
				cell = append(cell, model)
			}
			matrix.Field[rowIdx][colIdx] = cell
		}
	}

	return matrix, nil
}

func (r *Repository) CompleteSudokuMatrix(ctx context.Context, user models.User, matrixType models.SudokuMatrixType) error {
	today := time.Now().UTC().Format(time.DateOnly)
	key := rowKey(idString(user.Id), string(matrixType), today)
	_, err := r.collection("sudoku_matrix_completions").UpdateOne(
		ctx,
		bson.M{"_id": key},
		bson.M{"$setOnInsert": bson.M{
			"_id":          key,
			"user_id":      idString(user.Id),
			"matrix_type":  string(matrixType),
			"completed_on": today,
		}},
		options.Update().SetUpsert(true),
	)
	return err
}

func (r *Repository) GetSudokuCompletionStats(ctx context.Context, user models.User) (models.SudokuCompletionStats, error) {
	var completions []sudokuCompletionDoc
	if err := r.findAll(ctx, "sudoku_matrix_completions", bson.M{"user_id": idString(user.Id)}, &completions); err != nil {
		return models.SudokuCompletionStats{}, err
	}

	stats := models.SudokuCompletionStats{}
	for _, completion := range completions {
		switch models.SudokuMatrixType(completion.MatrixType) {
		case models.SudokuMatrixTypeDrivers:
			stats.DriverMatrices++
		case models.SudokuMatrixTypeTeams:
			stats.TeamMatrices++
		}
	}

	return stats, nil
}
