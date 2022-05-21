package repo

import (
	"context"
	"database/sql"

	"github.com/bogach-ivan/wb_assistant_be/pb"
)

type Authorization interface {
	CreateUser(ctx context.Context, username, email, password string) (*pb.CreateUserResponse, error)
	GetUserID(email, password string) (int, error)
	GetDetails(userId int32) (*pb.User, error)
	Update(in *pb.UpdateRequest) (*pb.UpdateResponse, error)
	InsertEmailConfirmToken(ctx context.Context, in *pb.InsertEmailConfirmTokenRequest) (*pb.InsertEmailConfirmTokenResponse, error)
	CheckAndDelEmailConfirmToken(ctx context.Context, in *pb.CheckAndDelEmailConfirmTokenRequest) (*pb.CheckAndDelEmailConfirmTokenResponse, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMySQL(db),
	}
}
