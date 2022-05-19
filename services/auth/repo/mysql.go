package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersTable            = "users"
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
	// user7:s$cret@tcp(127.0.0.1:3306)/testdb
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.DBName)
	fmt.Println(s)
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
