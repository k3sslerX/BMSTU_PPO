package models

type Team struct {
	Id      Uuid   `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
