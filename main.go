package main

import (
	"exam_service/pkg/handlers"
	"exam_service/pkg/logging"
	"github.com/spf13/viper"
	"log"
)

func init() {
	//logging
	logging.Logging()

	//app configuration
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("Fatal error config file: ", err)
	}
}

func main() {
	log.Println("Server is starting")
	handlers.Start()
}
