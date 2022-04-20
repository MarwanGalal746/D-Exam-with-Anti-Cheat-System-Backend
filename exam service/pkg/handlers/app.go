package handlers

import (
	"exam_service/pkg/domain/exam"
	"exam_service/pkg/driver"
	"exam_service/pkg/service"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	dbConnection := driver.GetDbConnection()

	examHandler := ExamHandlers{service.NewExamService(exam.NewExamRepositoryDb(dbConnection))}

	router.POST("/api/exam/create-exam", examHandler.Create)
	router.Run(":8888")

}
