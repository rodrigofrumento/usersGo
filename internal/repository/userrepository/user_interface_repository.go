package userrepository

import (
	"database/sql"

	"github.com/rodrigofrumento/usersGo/internal/database/sqlc"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser() error
}

func (r *repository) CreateUser() error {
	return nil
}
