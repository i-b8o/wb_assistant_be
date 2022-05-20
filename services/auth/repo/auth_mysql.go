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
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, expires, type) values (?, ?, ?, ?, ?)", usersTable)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}
	defer stmt.Close()

	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)
	res, err := stmt.ExecContext(ctx, user.Username, user.Email, user.Password, dt, "free")
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}

	fmt.Printf("ID: %d", id)
	return &pb.CreateUserResponse{ID: int32(id)}, nil
}

func (r *AuthMySQL) GetUserID(email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id  FROM %s WHERE email=? AND password=?", usersTable)
	err := r.db.QueryRow(query, email, password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthMySQL) GetDetails(userId int32) (*pb.User, error) {
	user := &pb.User{}
	query := fmt.Sprintf("SELECT id,username,email,password,expires,type FROM %s WHERE id=?", usersTable)
	err := r.db.QueryRow(query, userId).Scan(&user)
	if err != nil {
		return &pb.User{}, err
	}
	return user, nil
}
