package repo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable            = "users"
	verifiedsTable        = "verifieds"
	verifiedPasswordTable = "verifieds"
	resetPasswordTable    = "resets"
	actionsTable          = "actions"
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

	// Settings section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetConnMaxIdleTime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
