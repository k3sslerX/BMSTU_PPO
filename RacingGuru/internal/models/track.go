package models

import "github.com/google/uuid"

type Track struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Country string    `json:"country"`
	Length  int       `json:"length"`
	Turns   int       `json:"turns"`
}
