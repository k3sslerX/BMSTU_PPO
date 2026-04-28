package handlers

import (
	"RacingGuru/internal/core/auth"
	"RacingGuru/internal/logger"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Repo         Repo
	AdminSecrets *auth.AdminSecretManager
	Logger       *logger.Logger
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func NewHandler(repo Repo, logger *logger.Logger) *Handler {
	return &Handler{Repo: repo, AdminSecrets: auth.NewAdminSecretManager(), Logger: logger}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(h.requestLoggingMiddleware)
	r.Use(h.recovererMiddleware)

	r.HandleFunc("/", h.Info)
	r.Get("/swagger.yaml", h.SwaggerSpec)
	r.Get("/docs", h.SwaggerUI)

	r.Post("/login", h.Login)
	r.Post("/register", h.Register)
	r.Post("/setup/admin-secret", h.GenerateAdminSecret)
	r.Get("/drivers", h.ListDrivers)
	r.Get("/teams", h.ListTeams)

	r.Route("/stats", func(r chi.Router) {
		r.Get("/driver", h.StatsDriver)
		r.Get("/team", h.StatsTeam)
	})

	r.Group(func(r chi.Router) {
		r.Use(h.authMiddleware)
		r.Get("/me", h.Me)
		r.Post("/change-password", h.ChangePassword)
		r.Route("/users", func(r chi.Router) {
			r.Get("/favourites", h.ListFavourites)
			r.Post("/favourite-driver", h.ToggleFavouriteDriver)
			r.Post("/favourite-team", h.ToggleFavouriteTeam)
		})
		r.Route("/sudoku", func(r chi.Router) {
			r.Get("/drivers", h.SudokuDrivers)
			r.Get("/teams", h.SudokuTeams)
		})
		r.Route("/admin", func(r chi.Router) {
			r.Get("/cars", h.ListCars)
			r.Get("/car-participants", h.ListCarParticipants)
			r.Get("/championships", h.ListChampionships)
			r.Get("/races", h.ListRaces)
			r.Get("/tracks", h.ListTracks)
			r.Get("/users", h.ListUsers)
			r.Post("/drivers", h.CreateDriver)
			r.Patch("/drivers", h.UpdateDriver)
			r.Delete("/drivers/{id}", h.DeleteDriver)
			r.Patch("/users/role", h.UpdateUserRole)
			r.Post("/teams", h.CreateTeam)
			r.Patch("/teams", h.UpdateTeam)
			r.Delete("/teams/{id}", h.DeleteTeam)
			r.Post("/tracks", h.CreateTrack)
			r.Patch("/tracks", h.UpdateTrack)
			r.Delete("/tracks/{id}", h.DeleteTrack)
			r.Post("/races", h.CreateRace)
			r.Patch("/races", h.UpdateRace)
			r.Delete("/races/{id}", h.DeleteRace)
			r.Post("/car-participants", h.CreateCarParticipant)
			r.Patch("/car-participants", h.UpdateCarParticipant)
			r.Delete("/car-participants/{id}", h.DeleteCarParticipant)
			r.Patch("/car-participants/drivers", h.UpdateCarParticipantDrivers)
			r.Post("/race-results", h.UpsertRaceResult)
		})
	})

	return r
}

func (h *Handler) sendError(w http.ResponseWriter, message, code string, status int) {
	writeErrorResponse(w, message, code, status)
}

func writeErrorResponse(w http.ResponseWriter, message, code string, status int) {
	response := ErrorResponse{}
	response.Error.Code = code
	response.Error.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

// Info godoc
// @Summary Health check
// @Tags system
// @Success 200
// @Router / [get]
func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
