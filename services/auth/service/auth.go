package authservice

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
)

const salt = "jasfkldjasfkldjasklfs12-93234-0[23"

type AuthService struct {
	repo repo.AuthMySQL
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(repo repo.AuthMySQL) *AuthService {

	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {
	// Befor create hash password
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(ctx, user)
}

func (s *AuthService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return s.repo.GetUser(ctx, req)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
