package models

type Team struct {
	Id      uuid   `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
