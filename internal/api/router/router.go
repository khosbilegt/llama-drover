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

	r.Get("/herd", handlers.HandleListHerds)
	r.Get("/herd/{id}", handlers.HandleGetHerd)
	r.Post("/herd", handlers.HandleCreateHerd)
	r.Delete("/herd/{id}", handlers.HandleDeleteHerd)

	r.Get("/node", handlers.HandleListNodes)
	r.Get("/node/{id}", handlers.HandleFetchNode)
	r.Post("/node", handlers.HandleRegisterNode)
	r.Delete("/node/{id}", handlers.HandleDeleteNode)

	return r
}
