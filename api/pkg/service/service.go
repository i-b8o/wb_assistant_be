package service

import "github.com/bogach-ivan/wb_assistant_be/api"

type Authorization interface {
	CreateUser(user api.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService() *Service {
	return &Service{}
}
