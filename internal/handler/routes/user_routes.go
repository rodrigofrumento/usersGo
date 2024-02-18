package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rodrigofrumento/usersGo/internal/handler"
)

func InitUserRoutes(router chi.Router, h handler.Handler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Patch("/{id}", h.UpdateUser)
	})
}
