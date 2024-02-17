package userservice

import "github.com/rodrigofrumento/usersGo/internal/repository/userrepository"

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser() error
}

func (s *service) CreateUser() error {
	return nil
}
