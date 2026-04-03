package models

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Birthday    time.Time `json:"birthday"`
	Nationality string    `json:"nationality"`
}
