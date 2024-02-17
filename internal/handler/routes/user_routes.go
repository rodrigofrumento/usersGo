package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rodrigofrumento/usersGo/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
	})
}
