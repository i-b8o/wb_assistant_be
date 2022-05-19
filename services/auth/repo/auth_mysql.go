package repo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bogach-ivan/wb_assistant_be/api"
)

type AuthMySQL struct {
	db *sql.DB
}

func NewAuthMySQL(db *sql.DB) *AuthMySQL {

	return &AuthMySQL{
		db: db,
	}
}

func (r *AuthMySQL) CreateUser(user api.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password, expires, type) values ($1, $2, $3, $4, $5", usersTable)
	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, dt, "free")
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
