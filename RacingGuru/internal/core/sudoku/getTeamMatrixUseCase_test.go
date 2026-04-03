package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestGetTeamMatrixUseCase(t *testing.T) {
	var err error
	var matrix models.MatrixTeams

	testUc1 := NewGetTeamMatrixUseCase(&testRepo{returnError: false}, models.User{Role: models.RoleUser})
	matrix, err = testUc1.Run(context.Background())
	t.Log(formatTeamConditionMatrix(matrix))
	if err != nil {
		t.Error(err)
	}

	testUc2 := NewGetTeamMatrixUseCase(&testRepo{returnError: true}, models.User{Role: models.RoleUser})
	_, err = testUc2.Run(context.Background())
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}

func formatTeamConditionMatrix(matrix models.MatrixTeams) string {
	var sb strings.Builder

	sb.WriteString("Team condition matrix:\n")
	sb.WriteString(fmt.Sprintf("      | %-35s | %-35s | %-35s |\n",
		matrix.ConditionSpecs[0][0].Label,
		matrix.ConditionSpecs[0][1].Label,
		matrix.ConditionSpecs[0][2].Label,
	))

	for i := range matrix.ConditionSpecs[1] {
		sb.WriteString(fmt.Sprintf("%-5s | %-35s | %-35s | %-35s |\n",
			matrix.ConditionSpecs[1][i].Label,
			"x",
			"x",
			"x",
		))
	}

	return sb.String()
}
