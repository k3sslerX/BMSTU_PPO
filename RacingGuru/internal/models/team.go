package models

import "github.com/google/uuid"

type Team struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Country string    `json:"country"`
}
