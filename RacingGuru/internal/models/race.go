package models

import (
	"time"

	"github.com/google/uuid"
)

type Race struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	Type           int       `json:"type"`
	Duration       int       `json:"duration"`
	Track          Track     `json:"track"`
	ChampionshipId uuid.UUID `json:"championship_id"`
}
