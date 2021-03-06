package authservice

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "jasfkldjasfkldjasklfs12-93234-0[23"
	signingKey = "sadsadasdasdasdasdas"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repo.AuthMySQL
	pb.UnimplementedAuthServiceServer
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

func NewAuthService(repo repo.AuthMySQL) *AuthService {

	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	username := req.GetUsername()
	email := req.GetEmail()
	password := req.GetPassword()

	// Befor create hash password
	password = generatePasswordHash(password)

	return s.repo.CreateUser(ctx, username, email, password)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(ctx context.Context, in *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	id := s.repo.GetUserID(in.Email, generatePasswordHash(in.Password))
	if id == 0 {
		return &pb.GenerateTokenResponse{Token: ""}, nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt: time.Now().Unix()},
		int(id),
	})

	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return &pb.GenerateTokenResponse{}, err
	}
	return &pb.GenerateTokenResponse{Token: t}, nil
}

func (s *AuthService) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	token, err := jwt.ParseWithClaims(in.Token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return &pb.ParseTokenResponse{}, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return &pb.ParseTokenResponse{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return &pb.ParseTokenResponse{}, errors.New("token claims are not of type *tokenClaims")
	}
	return &pb.ParseTokenResponse{
		ID: int32(claims.UserId),
	}, nil
}

func (s *AuthService) GetDetails(ctx context.Context, in *pb.GetDetailsRequest) (*pb.User, error) {
	return s.repo.GetDetails(in.ID)
}

func (s *AuthService) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	if in.Password != "" {

		in.Password = generatePasswordHash(in.Password)

	}
	return s.repo.Update(in)
}

func (s *AuthService) InsertEmailConfirmToken(ctx context.Context, in *pb.InsertEmailConfirmTokenRequest) (*pb.InsertEmailConfirmTokenResponse, error) {
	return s.repo.InsertEmailConfirmToken(ctx, in)
}

func (s *AuthService) UpdateEmailVerificationToken(ctx context.Context, in *pb.UpdateEmailVerificationTokenRequest) (*pb.UpdateEmailVerificationTokenResponse, error) {
	password := generatePasswordHash(in.Password)
	return s.repo.UpdateEmailConfirmToken(in.Email, in.Token, password)
}

func (s *AuthService) CheckAndDelEmailConfirmToken(ctx context.Context, in *pb.CheckAndDelEmailConfirmTokenRequest) (*pb.CheckAndDelEmailConfirmTokenResponse, error) {
	return s.repo.CheckAndDelEmailConfirmToken(ctx, in)
}

func (s *AuthService) RecoverPassword(ctx context.Context, in *pb.RecoverPasswordRequest) (*pb.RecoverPasswordResponse, error) {
	password := generatePasswordHash(in.Password)

	err := s.repo.RecoverPassword(in.Email, password)
	if err != nil {
		return &pb.RecoverPasswordResponse{}, err
	}

	return &pb.RecoverPasswordResponse{}, nil
}

func (s *AuthService) Actions(ctx context.Context, in *pb.ActionsRequest) (*pb.ActionsResponse, error) {
	t, err := s.repo.Actions(ctx, in.ID, in.Action)
	if err != nil {
		return &pb.ActionsResponse{}, err
	}
	return &pb.ActionsResponse{Type: t}, nil
}
