package main

import (
	"D-Exam-with-Anti-Cheat-System-Backend/pkg/handlers"
	"github.com/spf13/viper"
	"log"
)

func init() {
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
	handlers.StartServer()
}
