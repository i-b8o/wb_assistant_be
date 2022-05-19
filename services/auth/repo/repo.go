package repo

import (
	"database/sql"

	"github.com/bogach-ivan/wb_assistant_be/api"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
	}
}
