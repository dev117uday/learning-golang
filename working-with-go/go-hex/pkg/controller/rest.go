package controller

import (
	"go-hex/pkg/service"

	"github.com/go-chi/chi"
)

func Handler(sr service.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/api/users", HandleAdding(sr))

	return router

}
