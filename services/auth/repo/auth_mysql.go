package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

func (r *AuthMySQL) CreateUser(ctx context.Context, username, email, password string) (*pb.CreateUserResponse, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, expires, type) values (?, ?, ?, ?, ?)", usersTable)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}
	defer stmt.Close()

	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)
	res, err := stmt.ExecContext(ctx, username, email, password, dt, "none")
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &pb.CreateUserResponse{ID: 0}, err
	}

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
	var days int32
	query := fmt.Sprintf("SELECT id,username,email,expires,type,DATEDIFF(CURDATE(), expires) FROM %s WHERE id=?", usersTable)
	err := r.db.QueryRow(query, userId).Scan(&user.ID, &user.Username, &user.Email, &user.Expires, &user.Type, &days)
	if err != nil {
		return &pb.User{}, err
	}
	// Expired?
	if days > 0 {
		user.Type = "end"
	}

	return user, nil
}

func (r *AuthMySQL) Update(in *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if in.Password != "" {
		setValues = append(setValues, "password=?")
		args = append(args, in.Password)
	}

	if in.Username != "" {
		setValues = append(setValues, "username=?")
		args = append(args, in.Username)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=?", usersTable, setQuery)

	args = append(args, in.ID)
	if _, err := r.db.Exec(query, args...); err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{}, nil
}

func (r *AuthMySQL) InsertEmailConfirmToken(ctx context.Context, in *pb.InsertEmailConfirmTokenRequest) (*pb.InsertEmailConfirmTokenResponse, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, token) values (?, ?)", verifiedsTable)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &pb.InsertEmailConfirmTokenResponse{}, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, in.ID, in.Token)
	if err != nil {
		return &pb.InsertEmailConfirmTokenResponse{}, err
	}

	return &pb.InsertEmailConfirmTokenResponse{}, nil
}

func (r *AuthMySQL) CheckAndDelEmailConfirmToken(ctx context.Context, in *pb.CheckAndDelEmailConfirmTokenRequest) (*pb.CheckAndDelEmailConfirmTokenResponse, error) {
	var id int32
	// Get user ID
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE token=?", verifiedsTable)
	err := r.db.QueryRow(query, in.Token).Scan(&id)
	if err != nil {
		return &pb.CheckAndDelEmailConfirmTokenResponse{}, err
	}
	// Delete row
	query = fmt.Sprintf("DELETE FROM %s WHERE user_id=?", verifiedsTable)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return &pb.CheckAndDelEmailConfirmTokenResponse{}, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return &pb.CheckAndDelEmailConfirmTokenResponse{}, err
	}
	// Update type
	query = fmt.Sprintf("UPDATE %s SET type='free' WHERE id=?", usersTable)

	if _, err := r.db.Exec(query, id); err != nil {
		return nil, err
	}
	return &pb.CheckAndDelEmailConfirmTokenResponse{}, nil
}
