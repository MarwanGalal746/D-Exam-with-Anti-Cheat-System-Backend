package logging

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	return logFile, nil
}

func Logging() {
	file, err := OpenLogFile("./logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = file
}
