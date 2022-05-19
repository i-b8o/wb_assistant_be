package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
	authservice "github.com/bogach-ivan/wb_assistant_be/services/auth/service"
	"github.com/i-rm/nonsense"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		nonsense.SendStringToTelegram("Server Panicking!")
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
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

var (
	auth_server_port string
	db_host          string
	db_username      string
	db_password      string
	db_name          string
)

func main() {
	defer recoveryFunction()
	go unixSig()

	port, err := strconv.Atoi(auth_server_port)
	if err != nil {
		logrus.Fatalf("cannot convert %s to int: %v", auth_server_port, err)

	}

	err = godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repo.NewMySQLDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	// store := authservice.NewDBStore(db_host, db_username, db_password, db_name)
	repo := repo.NewAuthMySQL(db)
	server := authservice.NewAuthService(*repo)
	pb.RegisterAuthServiceServer(grpcServer, server)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatal("cannot start server: ", err)
	}
	logrus.Printf("start server on address %s", address)

	err = grpcServer.Serve(listener)
	if err != nil {
		logrus.Fatal("cannot start server: ", err)
	}
}
