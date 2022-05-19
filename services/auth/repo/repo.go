package repo

import (
	"context"
	"database/sql"

	"github.com/bogach-ivan/wb_assistant_be/pb"
)

type Authorization interface {
	// CreateUser(user api.User) (int, error)
	CreateUser(context.Context, *pb.User) (*pb.CreateUserResponse, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
	}
}
