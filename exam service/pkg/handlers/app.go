package handlers

import (
	"exam_service/pkg/driver"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	dbConnection := driver.GetDbConnection()

	examHandler := Exa

	router.POST("/exam/create")
	router.Run(":8888")

}
