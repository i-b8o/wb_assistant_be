package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bogach-ivan/wb_assistant_be/pb"
)

type AuthMySQL struct {
	db *sql.DB
}

func NewAuthMySQL(db *sql.DB) *AuthMySQL {

	return &AuthMySQL{
		db: db,
	}
}

func (r *AuthMySQL) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, expires, type) values (?, ?, ?, ?, ?)", usersTable)

	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)
	fmt.Println(query, user.Username, user.Email, user.Password, dt, "free")
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, dt, "free")
	err := row.Scan(&id)
	if err != nil {
		return &pb.CreateUserResponse{Id: 0}, err
	}
	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
