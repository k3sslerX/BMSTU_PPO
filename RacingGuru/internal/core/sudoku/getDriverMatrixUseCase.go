package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetDriverMatrixUseCase struct {
	Repo Repo
	User models.User
}

func NewGetDriverMatrixUseCase(repo Repo, user models.User) *GetDriverMatrixUseCase {
	return &GetDriverMatrixUseCase{Repo: repo, User: user}
}

func (uc *GetDriverMatrixUseCase) Run(ctx context.Context) (models.MatrixDrivers, error) {
	matrix, err := generateConditionsDrivers()
	if err != nil {
		return models.MatrixDrivers{}, err
	}
	return uc.Repo.GetDriverMatrix(ctx, matrix)
}

func generateConditionsDrivers() (models.MatrixDrivers, error) {
	matrix := models.MatrixDrivers{
		Field:          [3][3][]models.Driver{},
		ConditionSpecs: [2][3]models.SudokuCondition{},
	}

	conditions := []models.SudokuCondition{
		{Label: "Больше 200 гонок", Field: "total_races", Op: models.ConditionOperatorGT, Value: 200},
		{Label: "Меньше 100 гонок", Field: "total_races", Op: models.ConditionOperatorLT, Value: 100},
		{Label: "Хотя бы 20 гонок", Field: "total_races", Op: models.ConditionOperatorGTE, Value: 20},
		{Label: "Больше 3 побед", Field: "total_wins", Op: models.ConditionOperatorGT, Value: 3},
		{Label: "Меньше 5 побед", Field: "total_wins", Op: models.ConditionOperatorLT, Value: 5},
		{Label: "Хотя бы 1 победа", Field: "total_wins", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без побед", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 3 подиумов", Field: "total_podiums", Op: models.ConditionOperatorGT, Value: 3},
		{Label: "Меньше 5 подиумов", Field: "total_podiums", Op: models.ConditionOperatorLT, Value: 5},
		{Label: "Хотя бы 1 подиум", Field: "total_podiums", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без подиумов", Field: "total_podiums", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 500 очков", Field: "total_points", Op: models.ConditionOperatorGT, Value: 500},
		{Label: "Меньше 500 очков", Field: "total_points", Op: models.ConditionOperatorLT, Value: 500},
		{Label: "Набирал очки", Field: "total_points", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Больше 5 поулов", Field: "total_poles", Op: models.ConditionOperatorGT, Value: 5},
		{Label: "Меньше 3 поулов", Field: "total_poles", Op: models.ConditionOperatorLT, Value: 3},
		{Label: "Хотя бы 1 поул", Field: "total_poles", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без поулов", Field: "total_poles", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Лучший финиш: победа", Field: "best_finish", Op: models.ConditionOperatorEQ, Value: 1},
		{Label: "Лучший финиш в топ-3", Field: "best_finish", Op: models.ConditionOperatorLTE, Value: 3},
		{Label: "Лучший финиш в топ-10", Field: "best_finish", Op: models.ConditionOperatorLTE, Value: 10},
		{Label: "Лучший финиш вне топ-3", Field: "best_finish", Op: models.ConditionOperatorGT, Value: 3},
		{Label: "Лучшая квалификация: поул", Field: "best_qualifying", Op: models.ConditionOperatorEQ, Value: 1},
		{Label: "Лучшая квалификация в топ-3", Field: "best_qualifying", Op: models.ConditionOperatorLTE, Value: 3},
		{Label: "Лучшая квалификация в топ-10", Field: "best_qualifying", Op: models.ConditionOperatorLTE, Value: 10},
		{Label: "Лучшая квалификация вне топ-3", Field: "best_qualifying", Op: models.ConditionOperatorGT, Value: 3},
		{Label: "Лучший результат в чемпионате в топ-5", Field: "best_championship_position", Op: models.ConditionOperatorLTE, Value: 5},
	}

	if err := fillMatrixConditions(conditions, &matrix.ConditionSpecs); err != nil {
		return models.MatrixDrivers{}, err
	}

	return matrix, nil
}
