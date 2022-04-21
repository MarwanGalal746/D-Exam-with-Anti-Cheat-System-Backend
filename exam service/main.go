package main

import (
	"exam_service/pkg/handlers"
	"exam_service/pkg/logging"
	"log"
)

func main() {

	//logging
	logging.Logging()

	log.Println("Server is starting")
	handlers.Start()
}
