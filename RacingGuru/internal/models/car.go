package models

import "github.com/google/uuid"

type Car struct {
	Id        uuid.UUID `json:"id"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	RaceClass string    `json:"raceclass,omitempty"`
}
