package models

type ConditionOperator string

const (
	ConditionOperatorGT  ConditionOperator = "gt"
	ConditionOperatorLT  ConditionOperator = "lt"
	ConditionOperatorGTE ConditionOperator = "gte"
	ConditionOperatorLTE ConditionOperator = "lte"
	ConditionOperatorEQ  ConditionOperator = "eq"
)

type SudokuCondition struct {
	Label string            `json:"label"`
	Field string            `json:"field"`
	Op    ConditionOperator `json:"op"`
	Value int               `json:"value"`
}

type MatrixDrivers struct {
	Field          [3][3][]Driver        `json:"field"`
	ConditionSpecs [2][3]SudokuCondition `json:"condition_specs"`
}

type MatrixTeams struct {
	Field          [3][3][]Team          `json:"field"`
	ConditionSpecs [2][3]SudokuCondition `json:"condition_specs"`
}

type SudokuMatrixType string

const (
	SudokuMatrixTypeDrivers SudokuMatrixType = "drivers"
	SudokuMatrixTypeTeams   SudokuMatrixType = "teams"
)

type SudokuCompletionStats struct {
	DriverMatrices int64 `json:"driver_matrices"`
	TeamMatrices   int64 `json:"team_matrices"`
}
