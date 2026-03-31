package models

type TeamStats struct {
	Team                     Team `json:"team"`
	TotalRaces               int  `json:"total_races"`
	TotalWins                int  `json:"total_wins"`
	TotalPodiums             int  `json:"total_podiums"`
	TotalPoints              int  `json:"total_points"`
	TotalPoles               int  `json:"total_poles"`
	BestFinish               int  `json:"best_finish"`
	BestQualifying           int  `json:"best_qualifying"`
	ChampionshipsWins        int  `json:"championships_wins"`
	BestChampionshipPosition int  `json:"best_championship_position"`
}
