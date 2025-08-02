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

	// Cluster routes
	r.Get("/cluster", handlers.HandleListClusters)
	r.Get("/cluster/{id}", handlers.HandleGetCluster)
	r.Post("/cluster", handlers.HandleCreateCluster)
	r.Delete("/cluster/{id}", handlers.HandleDeleteCluster)

	// Node routes
	r.Get("/node", handlers.HandleListNodes)
	r.Get("/node/{id}", handlers.HandleFetchNode)
	r.Post("/node", handlers.HandleCreateNode)
	r.Delete("/node/{id}", handlers.HandleDeleteNode)

	return r
}
