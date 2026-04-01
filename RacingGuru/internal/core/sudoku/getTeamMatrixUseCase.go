package sudoku

import (
	"RacingGuru/internal/models"
	"context"
)

type GetTeamMatrixUseCase struct {
	Repo Repo
	User models.User
}

func NewGetTeamMatrixUseCase(repo Repo, user models.User) *GetTeamMatrixUseCase {
	return &GetTeamMatrixUseCase{Repo: repo, User: user}
}

func (uc *GetTeamMatrixUseCase) Run(ctx context.Context) (models.MatrixTeams, error) {
	matrix, err := generateConditionsTeams()
	if err != nil {
		return models.MatrixTeams{}, err
	}
	return uc.Repo.GetTeamMatrix(ctx, matrix)
}

func generateConditionsTeams() (models.MatrixTeams, error) {
	matrix := models.MatrixTeams{
		Field:          [3][3][]models.Team{},
		ConditionSpecs: [2][3]models.SudokuCondition{},
	}

	conditions := []models.SudokuCondition{
		{Label: "Больше 300 гонок", Field: "total_races", Op: models.ConditionOperatorGT, Value: 300},
		{Label: "Меньше 100 гонок", Field: "total_races", Op: models.ConditionOperatorLT, Value: 100},
		{Label: "Хотя бы 1 гонка", Field: "total_races", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Больше 30 побед", Field: "total_wins", Op: models.ConditionOperatorGT, Value: 30},
		{Label: "Меньше 10 побед", Field: "total_wins", Op: models.ConditionOperatorLT, Value: 10},
		{Label: "Хотя бы 1 победа", Field: "total_wins", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без побед", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 75 подиумов", Field: "total_podiums", Op: models.ConditionOperatorGT, Value: 75},
		{Label: "Меньше 20 подиумов", Field: "total_podiums", Op: models.ConditionOperatorLT, Value: 20},
		{Label: "Хотя бы 1 подиум", Field: "total_podiums", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без подиумов", Field: "total_podiums", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 1000 очков", Field: "total_points", Op: models.ConditionOperatorGT, Value: 1000},
		{Label: "Меньше 250 очков", Field: "total_points", Op: models.ConditionOperatorLT, Value: 250},
		{Label: "Набирала очки", Field: "total_points", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без очков", Field: "total_points", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 20 поулов", Field: "total_poles", Op: models.ConditionOperatorGT, Value: 20},
		{Label: "Меньше 5 поулов", Field: "total_poles", Op: models.ConditionOperatorLT, Value: 5},
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
		{Label: "Лучший результат в чемпионате в топ-3", Field: "best_championship_position", Op: models.ConditionOperatorLTE, Value: 3},
	}

	if err := fillMatrixConditions(conditions, &matrix.ConditionSpecs); err != nil {
		return models.MatrixTeams{}, err
	}

	return matrix, nil
}
