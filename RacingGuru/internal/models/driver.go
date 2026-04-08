package models

import "github.com/google/uuid"

type Driver struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Birthday    string    `json:"birthday"`
	Nationality string    `json:"nationality"`
}
