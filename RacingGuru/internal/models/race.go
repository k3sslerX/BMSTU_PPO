package models

import (
	"time"

	"github.com/google/uuid"
)

type Race struct {
	Id       uuid.UUID     `json:"id"`
	Name     string        `json:"name"`
	Track    Track         `json:"track"`
	Date     time.Time     `json:"date"`
	Type     int           `json:"type"`
	Duration time.Duration `json:"duration"`
}
