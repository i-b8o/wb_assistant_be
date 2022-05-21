package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable            = "users"
	verifiedsTable        = "verifieds"
	verifiedPasswordTable = "verifieds"
	resetPasswordTable    = "resets"
)

type Config struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func NewMySQLDB(cfg Config) (*sql.DB, error) {

	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.DBName)

	db, err := sql.Open("mysql", s)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
