package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bogach-ivan/nonsense"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
	authservice "github.com/bogach-ivan/wb_assistant_be/services/auth/service"
	"github.com/spf13/viper"

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
	err := initConfig()
	port := os.Getenv("AUTH_SERVICE_PORT")

	db, err := repo.NewMySQLDB(repo.Config{
		Host:     os.Getenv("AUTH_SERVICE_DB_HOST"),
		Username: os.Getenv("AUTH_SERVICE_DB_USERNAME"),
		Password: os.Getenv("AUTH_SERVICE_DB_PASSWORD"),
		DBName:   os.Getenv("AUTH_SERVICE_DB_DBNAME"),
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

	address := fmt.Sprintf("0.0.0.0:%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatal("cannot start server: ", err)
	}
	logrus.Printf("start server on address %s", address)

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			logrus.Fatal("cannot start server: ", err)
		}
	}()
	logrus.Print("WB Assistant GRPC Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("WB Assistant GRPC Server Shutting Down")

	grpcServer.Stop()

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
