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
	redisDb, redisJsonDb := driver.GetDbConnection()

	examHandler := ExamHandlers{service.NewExamService(exam.NewExamRepositoryDb(redisDb, redisJsonDb))}

	router.POST("/api/exam/create-exam", examHandler.Create)
	router.GET("/api/exam/get-all-exams/:courseId", examHandler.GetCourseExams)
	router.GET("/api/exam/get-exam/:examId", examHandler.GetExam)

	router.Run(viper.GetString("SERVER_PORT"))

}
