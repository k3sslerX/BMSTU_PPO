package models

type Role string

const RoleAdmin Role = "admin"
const RoleUser Role = "user"

type User struct {
	Id               Uuid     `json:"id"`
	Name             string   `json:"name"`
	Password         string   `json:"password"`
	Role             Role     `json:"role"`
	FavouriteDrivers []Driver `json:"favourite_driers"`
	FavouriteTeams   []Team   `json:"favourite_teams"`
}
