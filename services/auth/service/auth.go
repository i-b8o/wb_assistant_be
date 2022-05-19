package authservice

import (
	"crypto/sha1"
	"fmt"

	"github.com/bogach-ivan/wb_assistant_be/api"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
)

const salt = "jasfkldjasfkldjasklfs12-93234-0[23"

type AuthService struct {
	repo repo.AuthMySQL
}

func NewAuthService(repo repo.AuthMySQL) *AuthService {

	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user api.User) (int, error) {
	// Befor create hash password
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
