package handlers

import "RacingGuru/internal/models"

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

type RegisterRequest struct {
	Name     string `json:"name" example:"Max Verstappen"`
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
	Secret   string `json:"secret,omitempty" example:"WQ3mYvB1d8jK2s4Qn7LpT0xZc5RaE9Uf"`
}

type AuthUserResponse struct {
	Name  string      `json:"name" example:"Max Verstappen"`
	Email string      `json:"email" example:"user@example.com"`
	Role  models.Role `json:"role" enums:"admin,user" example:"user"`
}

type GenerateAdminSecretResponse struct {
	Secret string `json:"secret" example:"WQ3mYvB1d8jK2s4Qn7LpT0xZc5RaE9Uf"`
}

type UpdateUserRoleRequest struct {
	Id   string      `json:"id" example:"77777777-7777-7777-7777-777777777777"`
	Role models.Role `json:"role" enums:"admin,user" example:"admin"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" example:"newSecret123"`
}

type FavouritesResponse struct {
	FavouriteDrivers []models.Driver `json:"favourite_drivers"`
	FavouriteTeams   []models.Team   `json:"favourite_teams"`
}

type DriverStatsResponse struct {
	Stats models.DriverStats `json:"driver_stats"`
}

type TeamStatsResponse struct {
	Stats models.TeamStats `json:"team_stats"`
}
