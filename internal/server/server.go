package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ss49919201/diary-article/internal/dicontainer"
	"github.com/ss49919201/diary-article/internal/handler"
)

func NewServer(addr string) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		dicontainer.MustInvoke[handler.HealthCheck]().Do(w, r)
	})

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
