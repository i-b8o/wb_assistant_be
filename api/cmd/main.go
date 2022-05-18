package main

import (
	"log"

	"github.com/bogach-ivan/wb_assistant_be/api"
	"github.com/bogach-ivan/wb_assistant_be/api/pkg/handler"
	"github.com/bogach-ivan/wb_assistant_be/api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	services := service.NewService()
	handlers := handler.NewHandler(services)
	// init server instance
	srv := new(api.Server)
	// run server
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
