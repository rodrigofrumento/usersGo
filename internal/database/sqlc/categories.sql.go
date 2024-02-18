// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: categories.sql

package sqlc

import (
	"context"
	"time"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO category (id, title, created_at, updated_at) 
VALUES ($1, $2, $3, $4)
`

type CreateCategoryParams struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory,
		arg.ID,
		arg.Title,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}