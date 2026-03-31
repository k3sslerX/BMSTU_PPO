package sudoku

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"testing"
)

func TestGetDriverMatrixUseCase(t *testing.T) {
	var err error

	testUc1 := NewGetDriverMatrixUseCase(&testRepo{returnError: false}, models.User{Role: models.RoleUser})
	_, err = testUc1.Run(context.Background())
	if err != nil {
		t.Error(err)
	}

	testUc2 := NewGetDriverMatrixUseCase(&testRepo{returnError: true}, models.User{Role: models.RoleUser})
	_, err = testUc2.Run(context.Background())
	if !errors.Is(err, shared.ErrorNotFound) {
		t.Error(err)
	}
}
