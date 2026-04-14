package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Repo Repo
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func NewHandler(repo Repo) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.HandleFunc("/", h.Info)
	r.Get("/swagger.yaml", h.SwaggerSpec)
	r.Get("/docs", h.SwaggerUI)

	r.Get("/login", h.Login)
	r.Post("/register", h.Register)

	r.Route("/stats", func(r chi.Router) {
		r.Get("/driver", h.StatsDriver)
		r.Get("/team", h.StatsTeam)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Post("/change-password", h.ChangePassword)
		r.Route("/users", func(r chi.Router) {
			r.Post("/favourite-driver", h.ToggleFavouriteDriver)
			r.Post("/favourite-team", h.ToggleFavouriteTeam)
		})
		r.Route("/sudoku", func(r chi.Router) {
			r.Get("/drivers", h.SudokuDrivers)
			r.Get("/teams", h.SudokuTeams)
		})
		r.Route("/admin", func(r chi.Router) {
			r.Post("/drivers", h.CreateDriver)
			r.Patch("/drivers", h.UpdateDriver)
			r.Patch("/users/role", h.UpdateUserRole)
			r.Post("/teams", h.CreateTeam)
			r.Patch("/teams", h.UpdateTeam)
			r.Post("/tracks", h.CreateTrack)
			r.Patch("/tracks", h.UpdateTrack)
			r.Post("/races", h.CreateRace)
			r.Patch("/races", h.UpdateRace)
			r.Post("/car-participants", h.CreateCarParticipant)
			r.Patch("/car-participants", h.UpdateCarParticipant)
			r.Patch("/car-participants/drivers", h.UpdateCarParticipantDrivers)
			r.Post("/race-results", h.UpsertRaceResult)
		})
	})

	return r
}

func (h *Handler) sendError(w http.ResponseWriter, message, code string, status int) {
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
