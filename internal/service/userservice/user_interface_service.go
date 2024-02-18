package userservice

import (
	"context"

	"github.com/rodrigofrumento/usersGo/internal/dto"
	"github.com/rodrigofrumento/usersGo/internal/handler/response"
	"github.com/rodrigofrumento/usersGo/internal/repository/userrepository"
)

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error
	Login(ctx context.Context, u dto.LoginDTO) (*response.UserAuthToken, error)
}
