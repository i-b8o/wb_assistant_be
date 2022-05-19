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

// TODO perform setting of db connection timeout and so on
func (r *AuthMySQL) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	// var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, expires, type) values (?, ?, ?, ?, ?)", usersTable)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return &pb.CreateUserResponse{Id: 0}, err
	}
	defer stmt.Close()

	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)
	res, err := stmt.ExecContext(ctx, user.Username, user.Email, user.Password, dt, "free")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return &pb.CreateUserResponse{Id: 0}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return &pb.CreateUserResponse{Id: 0}, err
	}

	fmt.Printf("ID: %d", id)
	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
