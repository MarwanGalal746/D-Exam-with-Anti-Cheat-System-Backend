package handlers

import (
	"exam_service/pkg/domain/exam"
	"exam_service/pkg/driver"
	"exam_service/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start() {
	router := gin.Default()
	dbConnection := driver.GetDbConnection()

	examHandler := ExamHandlers{service.NewExamService(exam.NewExamRepositoryDb(dbConnection))}

	router.POST("/api/exam/create-exam", examHandler.Create)
	//router.GET("/api/exam/get-all-exams", examHandler.GetAll)
	//router.GET("/api/exam/get-exam/:name", examHandler.GetExam)

	router.Run(viper.GetString("SERVER_PORT"))

}
