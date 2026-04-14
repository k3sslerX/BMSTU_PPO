package handlers

import "RacingGuru/internal/models"

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

type RegisterRequest struct {
	Name     string      `json:"name" example:"Max Verstappen"`
	Email    string      `json:"email" example:"user@example.com"`
	Password string      `json:"password" example:"secret123"`
	Role     models.Role `json:"role" enums:"admin,user" example:"user"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" example:"newSecret123"`
}

type DriverStatsResponse struct {
	Stats models.DriverStats `json:"driver_stats"`
}

type TeamStatsResponse struct {
	Stats models.TeamStats `json:"team_stats"`
}
