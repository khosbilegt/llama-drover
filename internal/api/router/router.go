package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/khosbilegt/llama-drover/internal/api/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// r.Get("/health", handlers.HandleHealthCheck)
	r.Post("/prompt", handlers.HandleGeneratePrompt)

	return r
}
