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

	r.Post("/login", h.Login)
	r.Post("/register", h.Register)

	r.Route("/stats", func(r chi.Router) {
		r.Get("/driver", h.StatsDriver)
		r.Get("/team", h.StatsTeam)
	})

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

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

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
