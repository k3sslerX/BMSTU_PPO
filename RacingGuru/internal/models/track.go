package models

type Track struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Length  int    `json:"length"`
	Turns   int    `json:"turns"`
}
