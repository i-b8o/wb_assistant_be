package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bogach-ivan/wb_assistant_be"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/pkg/handler"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

// TODO Perform telgram notification about erros
// @title WB Assistant API
// @version 1.0
// @description API Server for WB Assistant

// @host 188.93.210.165:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// GRPC client creation
	addrAuth := fmt.Sprintf("%s:%s", os.Getenv("AUTH_SERVICE_IP"), os.Getenv("AUTH_SERVICE_PORT"))
	authClientConn, err := grpc.Dial(addrAuth, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("error creating auth grpc connection: %s", err.Error())
	}
	authClient := pb.NewAuthServiceClient(authClientConn)

	addrMail := fmt.Sprintf("%s:%s", os.Getenv("MAIL_SERVICE_IP"), os.Getenv("MAIL_SERVICE_PORT"))
	mailClientConn, err := grpc.Dial(addrMail, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("error creating auth grpc connection: %s", err.Error())
	}
	mailClient := pb.NewMailServiceClient(mailClientConn)

	handlers := handler.NewHandler(authClient, mailClient)
	// init server instance
	srv := new(wb_assistant_be.Server)
	// run server
	go func() {
		if err := srv.Run(os.Getenv("API_SERVICE_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("WB Assistant started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("WB Assistant Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
