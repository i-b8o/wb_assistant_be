package mailservice

import (
	"context"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server ...
type Server struct {
	post Post

	// pb.UnimplementedAuthServiceServer
	pb.UnimplementedMailServiceServer
}

// NewServer ...
func NewServer(post Post) *Server {
	return &Server{
		post:                           post,
		UnimplementedMailServiceServer: pb.UnimplementedMailServiceServer{},
	}
}

func (server *Server) Confirm(ctx context.Context, req *pb.MailConfirmRequest) (*pb.MailConfirmResponse, error) {
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
	resp := server.post.Confirm(req.Url, req.Email, req.Pass)
	return resp, nil
}

func (server *Server) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.ResetResponse, error) {
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
	resp := server.post.Reset(req.Email, req.Password)
	return resp, nil
}
