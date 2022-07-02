package handlers

import (
	repositories2 "exam_service/domain/repositories"
	"exam_service/driver"
	"exam_service/messaging"
	service2 "exam_service/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"time"
)

var validate *validator.Validate

func Start() {

	validate = validator.New()

	router := gin.Default()
	redisDb, redisJsonDb, sqlDb := driver.GetDbConnection()

	examHandler := ExamHandlers{
		service2.NewExamService(repositories2.NewExamRepositoryDb(redisDb, redisJsonDb))}
	questionHandler := QuestionHandlers{
		service2.NewQuestionService(repositories2.NewQuestionRepositoryDb(redisDb, redisJsonDb))}
	studentGradeHandler := StudentGradeHandlers{
		service2.NewStudentGradeService(repositories2.NewStudentGradeRepositoryDb(sqlDb, redisDb, redisJsonDb))}

	go messaging.DeleteCourseExams(repositories2.NewExamRepositoryDb(redisDb, redisJsonDb))

	//enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	//exam endpoints
	router.POST("/api/exam/create-exam", examHandler.Create)
	router.POST("/api/exam/get-all-exams", examHandler.GetCourseExams)
	router.GET("/api/exam/get-exam/:examId", examHandler.GetExam)
	router.DELETE("/api/exam/delete-exam/:examId", examHandler.DelExam)
	router.PUT("/api/exam/update-exam-info/:examId", examHandler.UpdateExamInfo)

	//question endpoints
	router.POST("/api/exam/add-question/:examId", questionHandler.Add)
	router.DELETE("/api/exam/delete-question/:examId/:questionId", questionHandler.Delete)
	router.PUT("/api/exam/update-question/:examId/:questionId", questionHandler.Update)

	//student grade endpoints
	router.POST("/api/exam/add-student-grade/:userId/:courseId/:examId", studentGradeHandler.Add)
	router.GET("/api/exam/get-all-student-grades/:userId", studentGradeHandler.GetAllStudentGrades)
	router.GET("/api/exam/get-all-course-grades/:courseId", studentGradeHandler.GetAllCourseGrades)
	router.GET("/api/exam/get-all-exam-grades/:examId", studentGradeHandler.GetAllExamGrades)
	router.GET("/api/exam/get-user-exam-grade/:userId/:courseId/:examId",
		studentGradeHandler.GetUserCourseExamGrade)
	router.DELETE("/api/exam/delete-all-student-grades/:userId",
		studentGradeHandler.DeleteAllStudentGrades)
	router.DELETE("/api/exam/delete-all-course-grades/:courseId",
		studentGradeHandler.DeleteAllCourseGrades)
	router.DELETE("/api/exam/delete-all-exam-grades/:examId",
		studentGradeHandler.DeleteAllExamGrades)
	router.DELETE("/api/exam/delete-user-exam-grade/:userId/:courseId/:examId",
		studentGradeHandler.DeleteUserCourseExamGrade)

	router.Run(viper.GetString("SERVER_PORT"))

}
