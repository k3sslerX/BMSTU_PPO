package models

import "github.com/google/uuid"

type Championship struct {
	Id   uuid.UUID `json:"id"`
	Year int       `json:"year"`
}
