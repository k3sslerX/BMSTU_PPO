package models

import "github.com/google/uuid"

type CarParticipant struct {
	Id      uuid.UUID `json:"id"`
	CarID   uuid.UUID `json:"car_id"`
	TeamID  uuid.UUID `json:"team_id"`
	Number  string    `json:"number"`
	Drivers []Driver  `json:"drivers"`
}
