package sudoku

import (
	"RacingGuru/internal/models"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func fillMatrixConditions(
	conditions []models.SudokuCondition,
	specs *[2][3]models.SudokuCondition,
) error {
	rowCount := len(specs[0])
	colCount := len(specs[1])
	requiredConditions := rowCount + colCount
	if len(conditions) < requiredConditions {
		return fmt.Errorf("not enough sudoku conditions: got %d, need %d", len(conditions), requiredConditions)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(conditions), func(i, j int) {
		conditions[i], conditions[j] = conditions[j], conditions[i]
	})

	rowConditions, colConditions, ok := selectCompatibleConditions(conditions, rowCount, colCount)
	if !ok {
		return fmt.Errorf("failed to build compatible sudoku conditions")
	}

	for i := 0; i < rowCount; i++ {
		specs[0][i] = rowConditions[i]
	}
	for i := 0; i < colCount; i++ {
		specs[1][i] = colConditions[i]
	}

	return nil
}

func selectCompatibleConditions(
	conditions []models.SudokuCondition,
	rowCount int,
	colCount int,
) ([]models.SudokuCondition, []models.SudokuCondition, bool) {
	rows := make([]models.SudokuCondition, 0, rowCount)
	cols := make([]models.SudokuCondition, 0, colCount)
	used := make([]bool, len(conditions))

	var pickRows func(start int) bool
	var pickCols func(start int) bool

	pickRows = func(start int) bool {
		if len(rows) == rowCount {
			return pickCols(0)
		}

		for i := start; i < len(conditions); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			rows = append(rows, conditions[i])
			if pickRows(i + 1) {
				return true
			}
			rows = rows[:len(rows)-1]
			used[i] = false
		}
		return false
	}

	pickCols = func(start int) bool {
		if len(cols) == colCount {
			return true
		}

		for i := start; i < len(conditions); i++ {
			if used[i] {
				continue
			}

			candidate := conditions[i]
			compatible := true
			for _, rowCondition := range rows {
				if !conditionsCompatible(rowCondition, candidate) {
					compatible = false
					break
				}
			}
			if !compatible {
				continue
			}

			used[i] = true
			cols = append(cols, candidate)
			if pickCols(i + 1) {
				return true
			}
			cols = cols[:len(cols)-1]
			used[i] = false
		}
		return false
	}

	if !pickRows(0) {
		return nil, nil, false
	}
	return rows, cols, true
}

func conditionsCompatible(left models.SudokuCondition, right models.SudokuCondition) bool {
	if left.Field == right.Field {
		leftRange := conditionRange(left)
		rightRange := conditionRange(right)
		return rangesIntersect(leftRange, rightRange)
	}

	return crossFieldConditionsCompatible(left, right) && crossFieldConditionsCompatible(right, left)
}

func crossFieldConditionsCompatible(source models.SudokuCondition, target models.SudokuCondition) bool {
	switch source.Field {
	case "total_wins":
		if conditionRequiresAtLeast(source, 1) {
			if target.Field == "best_finish" {
				return rangesIntersect(conditionRange(target), valueRangeForExact(1))
			}
			if target.Field == "total_podiums" {
				return rangesIntersect(conditionRange(target), valueRangeForAtLeast(1))
			}
		}
		if conditionRequiresExactly(source, 0) && target.Field == "best_finish" {
			return !rangesIntersect(conditionRange(target), valueRangeForExact(1))
		}
	case "total_podiums":
		if conditionRequiresAtLeast(source, 1) && target.Field == "best_finish" {
			return rangesIntersect(conditionRange(target), valueRangeForAtMost(3))
		}
		if conditionRequiresExactly(source, 0) && target.Field == "best_finish" {
			return !rangesIntersect(conditionRange(target), valueRangeForAtMost(3))
		}
	case "total_poles":
		if conditionRequiresAtLeast(source, 1) && target.Field == "best_qualifying" {
			return rangesIntersect(conditionRange(target), valueRangeForExact(1))
		}
		if conditionRequiresExactly(source, 0) && target.Field == "best_qualifying" {
			return !rangesIntersect(conditionRange(target), valueRangeForExact(1))
		}
	}

	return true
}

type valueRange struct {
	min          int
	max          int
	minInclusive bool
	maxInclusive bool
}

func conditionRange(condition models.SudokuCondition) valueRange {
	switch condition.Op {
	case models.ConditionOperatorGT:
		return valueRange{
			min:          condition.Value,
			max:          math.MaxInt,
			minInclusive: false,
			maxInclusive: true,
		}
	case models.ConditionOperatorGTE:
		return valueRange{
			min:          condition.Value,
			max:          math.MaxInt,
			minInclusive: true,
			maxInclusive: true,
		}
	case models.ConditionOperatorLT:
		return valueRange{
			min:          math.MinInt,
			max:          condition.Value,
			minInclusive: true,
			maxInclusive: false,
		}
	case models.ConditionOperatorLTE:
		return valueRange{
			min:          math.MinInt,
			max:          condition.Value,
			minInclusive: true,
			maxInclusive: true,
		}
	case models.ConditionOperatorEQ:
		return valueRange{
			min:          condition.Value,
			max:          condition.Value,
			minInclusive: true,
			maxInclusive: true,
		}
	default:
		return valueRange{
			min:          math.MinInt,
			max:          math.MaxInt,
			minInclusive: true,
			maxInclusive: true,
		}
	}
}

func valueRangeForExact(value int) valueRange {
	return valueRange{
		min:          value,
		max:          value,
		minInclusive: true,
		maxInclusive: true,
	}
}

func valueRangeForAtLeast(value int) valueRange {
	return valueRange{
		min:          value,
		max:          math.MaxInt,
		minInclusive: true,
		maxInclusive: true,
	}
}

func valueRangeForAtMost(value int) valueRange {
	return valueRange{
		min:          math.MinInt,
		max:          value,
		minInclusive: true,
		maxInclusive: true,
	}
}

func conditionRequiresAtLeast(condition models.SudokuCondition, minValue int) bool {
	conditionRange := conditionRange(condition)
	requiredRange := valueRangeForAtLeast(minValue)
	return rangeContains(conditionRange, requiredRange)
}

func conditionRequiresExactly(condition models.SudokuCondition, value int) bool {
	return rangeContains(conditionRange(condition), valueRangeForExact(value))
}

func rangeContains(container valueRange, subset valueRange) bool {
	return rangeCoversLowerBound(container, subset) && rangeCoversUpperBound(container, subset)
}

func rangeCoversLowerBound(container valueRange, subset valueRange) bool {
	if container.min < subset.min {
		return true
	}
	if container.min > subset.min {
		return false
	}
	return container.minInclusive || !subset.minInclusive
}

func rangeCoversUpperBound(container valueRange, subset valueRange) bool {
	if container.max > subset.max {
		return true
	}
	if container.max < subset.max {
		return false
	}
	return container.maxInclusive || !subset.maxInclusive
}

func rangesIntersect(left valueRange, right valueRange) bool {
	lower := left.min
	lowerInclusive := left.minInclusive
	if right.min > lower {
		lower = right.min
		lowerInclusive = right.minInclusive
	} else if right.min == lower {
		lowerInclusive = left.minInclusive && right.minInclusive
	}

	upper := left.max
	upperInclusive := left.maxInclusive
	if right.max < upper {
		upper = right.max
		upperInclusive = right.maxInclusive
	} else if right.max == upper {
		upperInclusive = left.maxInclusive && right.maxInclusive
	}

	if lower < upper {
		return true
	}
	if lower > upper {
		return false
	}
	return lowerInclusive && upperInclusive
}
