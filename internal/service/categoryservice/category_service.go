package categoryservice

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodrigofrumento/usersGo/internal/dto"
	"github.com/rodrigofrumento/usersGo/internal/entity"
)

func (s *service) CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error {
	categoryEntity := entity.CategoryEntity{
		ID:        uuid.New().String(),
		Title:     u.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.CreateCategory(ctx, &categoryEntity)
	if err != nil {
		return errors.New("error to create category")
	}
	return nil
}
