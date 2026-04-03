package models

type RaceResults struct {
	Race    *Race
	Team    *Team
	Drivers []*Driver
	RacePos int
	QualPos int
}
