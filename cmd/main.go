package main

import (
	"fmt"

	"github.com/bogach-ivan/wb_assistant_be"
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/pkg/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := initConfig()
	if err != nil {

		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	fmt.Println(viper.GetString("db.dbname"))
	// GRPC client creation
	addr := fmt.Sprintf("%s:%s", viper.GetString("grpc.ip"), viper.GetString("grpc.port"))
	clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("error creating grpc connection: %s", err.Error())
	}
	client := pb.NewAuthServiceClient(clientConn)

	handlers := handler.NewHandler(client)
	// init server instance
	srv := new(wb_assistant_be.Server)
	// run server
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
