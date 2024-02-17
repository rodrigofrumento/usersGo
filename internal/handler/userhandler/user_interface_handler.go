package userhandler

import (
	"net/http"

	"github.com/rodrigofrumento/usersGo/internal/service/userservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {}
