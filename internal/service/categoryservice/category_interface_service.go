package categoryservice

import (
	"context"

	"github.com/rodrigofrumento/usersGo/internal/dto"
	"github.com/rodrigofrumento/usersGo/internal/repository/categoryrepository"
)

func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
	return &service{
		repo,
	}
}

type service struct {
	repo categoryrepository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
}
