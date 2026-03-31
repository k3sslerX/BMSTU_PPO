package models

type MatrixDrivers struct {
	Field      [3][3][]Driver `json:"field"`
	Conditions [2][3]string   `json:"conditions"`
}

type MatrixTeams struct {
	Field      [3][3][]Team `json:"field"`
	Conditions [2][3]string `json:"conditions"`
}

type MatrixTracks struct {
	Field      [3][3][]Track `json:"field"`
	Conditions [2][3]string  `json:"conditions"`
}
