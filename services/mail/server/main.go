package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bogach-ivan/nonsense"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	mailservice "github.com/bogach-ivan/wb_assistant_be/services/mail/service"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		nonsense.SendStringToTelegram("Server Panicking!")
		logrus.Println(recoveryMessage)
	}
	logrus.Println("This is recovery function...")
}

func unixSig() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for sig := range ch {
		switch sig {
		case syscall.SIGHUP: //Abort signal from abort(3)
			nonsense.SendStringToTelegram("sigterm received: " + "SIGHUP")
			os.Exit(0)
		case syscall.SIGINT: //Abort signal from abort(3)
			nonsense.SendStringToTelegram("sigterm received: " + "SIGINT")
			os.Exit(0)
		case syscall.SIGTERM: //Abort signal from abort(3)
			nonsense.SendStringToTelegram("sigterm received: " + "SIGTERM")
			os.Exit(0)
		case syscall.SIGQUIT: //Abort signal from abort(3)
			nonsense.SendStringToTelegram("sigterm received: " + "SIGQUIT")
			os.Exit(0)
		}
	}
}

func main() {
	defer recoveryFunction()
	go unixSig()
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// GRPC client creation
	grpcServer := grpc.NewServer()

	store := mailservice.NewPostFix()
	server := mailservice.NewServer(store)
	pb.RegisterMailServiceServer(grpcServer, server)

	address := fmt.Sprintf("0.0.0.0:%s", "1982")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	logrus.Printf("start server on address %s", address)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
