package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Config struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func NewMySQLDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.DBName))
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
