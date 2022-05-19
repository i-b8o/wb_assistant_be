package authservice

import (
	"context"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server ...
type Server struct {
	store Store

	pb.UnimplementedAuthServiceServer
}

// NewServer ...
func NewServer(store Store) *Server {
	return &Server{
		store:                          store,
		UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
	}
}

func (server *Server) SetTokenToPassReset(ctx context.Context, req *pb.SetTokenToPassResetRequest) (*pb.SetTokenToPassResetResponse, error) {
	// check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.SetTokenToPassReset(req.Email, req.Token)
	return resp, nil
}
func (server *Server) PassReset(ctx context.Context, req *pb.PassResetRequest) (*pb.PassResetResponse, error) {
	// / check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.PassReset(req.Email, req.NewPass, req.Token)
	return resp, nil
}

func (server *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.Register(req.Username, req.Email, req.Pass, req.Token)

	return resp, nil
}

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.Login(req.Email, req.Password)
	return resp, nil

}

func (server *Server) UpdateVerificationToken(ctx context.Context, req *pb.UpdateVerificationTokenRequest) (*pb.UpdateVerificationTokenResponse, error) {
	// check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.UpdateVerificationToken(req.Email, req.Pass, req.Token)
	return resp, nil

}

func (server *Server) Confirm(ctx context.Context, req *pb.AuthConfirmRequest) (*pb.AuthConfirmResponse, error) {
	// check if the request is cancelled by the client
	if ctx.Err() == context.Canceled {
		logrus.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	// check if the request is timeout
	if ctx.Err() == context.DeadlineExceeded {
		logrus.Print("deadline is exceed")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	resp := server.store.Confirm(req.Token)
	return resp, nil
}
