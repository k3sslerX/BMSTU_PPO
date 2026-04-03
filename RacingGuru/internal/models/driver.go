package models

type Driver struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Birthday    string `json:"birthday"`
	Nationality string `json:"nationality"`
}
