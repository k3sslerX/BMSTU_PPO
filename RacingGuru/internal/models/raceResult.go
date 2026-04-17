package models

import "github.com/google/uuid"

type RaceResult struct {
	RaceID           uuid.UUID `json:"race_id"`
	CarParticipantID uuid.UUID `json:"car_participant_id"`
	FinishPos        int       `json:"finish_pos"`
	QualifyingPos    int       `json:"qualifying_pos"`
}
