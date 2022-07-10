package main

import (
	"exam_service/pkg/handlers"
	"exam_service/pkg/logging"
	"log"
)

func init() {
	//logging
	logging.Logging()
}

func main() {
	log.Println("Server is starting")
	handlers.Start()
}
