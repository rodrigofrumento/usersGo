package handler

import (
	"net/http"

	"github.com/rodrigofrumento/usersGo/internal/service/userservice"
)

func NewHandler(userService userservice.UserService) Handler {
	return &handler{
		userService: userService,
	}
}

type handler struct {
	userService userservice.UserService
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
