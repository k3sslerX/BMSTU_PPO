package sudoku

import (
	"RacingGuru/internal/models"
	"errors"
	"testing"
)

func TestFillMatrixConditionsAvoidsConflictingPairs(t *testing.T) {
	conditions := []models.SudokuCondition{
		{Label: "Лучший финиш в топ-3", Field: "best_finish", Op: models.ConditionOperatorLTE, Value: 3},
		{Label: "Лучший финиш вне топ-5", Field: "best_finish", Op: models.ConditionOperatorGT, Value: 5},
		{Label: "Есть титул", Field: "championships_wins", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без титулов", Field: "championships_wins", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 200 гонок", Field: "total_races", Op: models.ConditionOperatorGT, Value: 200},
		{Label: "Больше 20 побед", Field: "total_wins", Op: models.ConditionOperatorGT, Value: 20},
	}

	var specs [2][3]models.SudokuCondition

	if err := fillMatrixConditions(conditions, &specs); err != nil {
		t.Fatalf("fillMatrixConditions returned error: %v", err)
	}

	for _, rowCondition := range specs[0] {
		for _, colCondition := range specs[1] {
			if !conditionsCompatible(rowCondition, colCondition) {
				t.Fatalf("found conflicting pair on matrix axes: %+v and %+v", rowCondition, colCondition)
			}
		}
	}
}

func TestFillMatrixConditionsReturnsErrorWhenNotEnoughConditions(t *testing.T) {
	conditions := []models.SudokuCondition{
		{Label: "Есть титул", Field: "championships_wins", Op: models.ConditionOperatorGTE, Value: 1},
		{Label: "Без титулов", Field: "championships_wins", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Больше 200 гонок", Field: "total_races", Op: models.ConditionOperatorGT, Value: 200},
	}

	var specs [2][3]models.SudokuCondition

	err := fillMatrixConditions(conditions, &specs)
	if err == nil {
		t.Fatal("expected error for insufficient conditions, got nil")
	}
}

func TestFillMatrixConditionsReturnsErrorWhenNoCompatibleLayoutExists(t *testing.T) {
	conditions := []models.SudokuCondition{
		{Label: "Ровно 0 побед", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 0},
		{Label: "Ровно 1 победа", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 1},
		{Label: "Ровно 2 победы", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 2},
		{Label: "Ровно 3 победы", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 3},
		{Label: "Ровно 4 победы", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 4},
		{Label: "Ровно 5 побед", Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 5},
	}

	var specs [2][3]models.SudokuCondition

	err := fillMatrixConditions(conditions, &specs)
	if err == nil {
		t.Fatal("expected error for incompatible conditions, got nil")
	}
}

func TestConditionsCompatible(t *testing.T) {
	tests := []struct {
		name     string
		left     models.SudokuCondition
		right    models.SudokuCondition
		expected bool
	}{
		{
			name: "different fields are compatible",
			left: models.SudokuCondition{Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 0},
			right: models.SudokuCondition{
				Field: "total_podiums", Op: models.ConditionOperatorEQ, Value: 0,
			},
			expected: true,
		},
		{
			name: "at least one win is incompatible with best finish outside top 3",
			left: models.SudokuCondition{Field: "total_wins", Op: models.ConditionOperatorGTE, Value: 1},
			right: models.SudokuCondition{
				Field: "best_finish", Op: models.ConditionOperatorGT, Value: 3,
			},
			expected: false,
		},
		{
			name: "no wins is incompatible with best finish win",
			left: models.SudokuCondition{Field: "total_wins", Op: models.ConditionOperatorEQ, Value: 0},
			right: models.SudokuCondition{
				Field: "best_finish", Op: models.ConditionOperatorEQ, Value: 1,
			},
			expected: false,
		},
		{
			name: "at least one podium is incompatible with best finish outside top 3",
			left: models.SudokuCondition{Field: "total_podiums", Op: models.ConditionOperatorGTE, Value: 1},
			right: models.SudokuCondition{
				Field: "best_finish", Op: models.ConditionOperatorGT, Value: 3,
			},
			expected: false,
		},
		{
			name: "at least one pole is incompatible with best qualifying outside top 3",
			left: models.SudokuCondition{Field: "total_poles", Op: models.ConditionOperatorGTE, Value: 1},
			right: models.SudokuCondition{
				Field: "best_qualifying", Op: models.ConditionOperatorGT, Value: 3,
			},
			expected: false,
		},
		{
			name: "overlapping same-field ranges are compatible",
			left: models.SudokuCondition{Field: "best_finish", Op: models.ConditionOperatorLTE, Value: 3},
			right: models.SudokuCondition{
				Field: "best_finish", Op: models.ConditionOperatorGT, Value: 1,
			},
			expected: true,
		},
		{
			name: "disjoint same-field ranges are incompatible",
			left: models.SudokuCondition{Field: "best_finish", Op: models.ConditionOperatorLTE, Value: 3},
			right: models.SudokuCondition{
				Field: "best_finish", Op: models.ConditionOperatorGT, Value: 5,
			},
			expected: false,
		},
		{
			name: "equal and gte same value are compatible",
			left: models.SudokuCondition{Field: "championships_wins", Op: models.ConditionOperatorEQ, Value: 1},
			right: models.SudokuCondition{
				Field: "championships_wins", Op: models.ConditionOperatorGTE, Value: 1,
			},
			expected: true,
		},
		{
			name: "equal and gt same value are incompatible",
			left: models.SudokuCondition{Field: "championships_wins", Op: models.ConditionOperatorEQ, Value: 1},
			right: models.SudokuCondition{
				Field: "championships_wins", Op: models.ConditionOperatorGT, Value: 1,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := conditionsCompatible(tt.left, tt.right)
			if actual != tt.expected {
				t.Fatalf("conditionsCompatible() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestGenerateConditionsDriversProducesUsableMatrix(t *testing.T) {
	matrix, err := generateConditionsDrivers()
	if err != nil {
		t.Fatalf("generateConditionsDrivers returned error: %v", err)
	}

	assertConditionMatrixValid(t, matrix.ConditionSpecs)
}

func TestGenerateConditionsTeamsProducesUsableMatrix(t *testing.T) {
	matrix, err := generateConditionsTeams()
	if err != nil {
		t.Fatalf("generateConditionsTeams returned error: %v", err)
	}

	assertConditionMatrixValid(t, matrix.ConditionSpecs)
}

func assertConditionMatrixValid(t *testing.T, specs [2][3]models.SudokuCondition) {
	t.Helper()

	seen := make(map[string]struct{}, 6)
	for axisIdx, axisConditions := range specs {
		for conditionIdx, condition := range axisConditions {
			if condition.Label == "" {
				t.Fatalf("empty label at axis %d condition %d", axisIdx, conditionIdx)
			}
			if condition.Field == "" {
				t.Fatalf("empty field at axis %d condition %d", axisIdx, conditionIdx)
			}
			key := condition.Label + "|" + condition.Field
			if _, exists := seen[key]; exists {
				t.Fatalf("duplicate condition selected: %+v", condition)
			}
			seen[key] = struct{}{}
		}
	}

	for _, rowCondition := range specs[0] {
		for _, colCondition := range specs[1] {
			if !conditionsCompatible(rowCondition, colCondition) {
				t.Fatalf("found conflicting pair on matrix axes: %+v and %+v", rowCondition, colCondition)
			}
		}
	}
}

func TestFillMatrixConditionsErrorMessagesAreMeaningful(t *testing.T) {
	var specs [2][3]models.SudokuCondition
	err := fillMatrixConditions(nil, &specs)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if errors.Is(err, nil) {
		t.Fatal("expected non-nil error identity")
	}
	if err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}
