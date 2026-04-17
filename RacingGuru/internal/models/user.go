package models

import "github.com/google/uuid"

type Role string

const RoleAdmin Role = "admin"
const RoleUser Role = "user"

type User struct {
	Id               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	Role             Role      `json:"role"`
	FavouriteDrivers []Driver  `json:"favourite_driers"`
	FavouriteTeams   []Team    `json:"favourite_teams"`
}
