package handlers

import (
	"exam_service/pkg/domain/repositories"
	"exam_service/pkg/driver"
	"exam_service/pkg/messaging"
	"exam_service/pkg/service"
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
		service.NewExamService(repositories.NewExamRepositoryDb(redisDb, redisJsonDb))}
	questionHandler := QuestionHandlers{
		service.NewQuestionService(repositories.NewQuestionRepositoryDb(redisDb, redisJsonDb))}
	studentGradeHandler := StudentGradeHandlers{
		service.NewStudentGradeService(repositories.NewStudentGradeRepositoryDb(sqlDb, redisDb, redisJsonDb))}
	examSubmissionHandler := ExamSubmissionHandlers{
		service.NewExamSubmissionService(repositories.NewExamSubmissionRepositoryDb(redisDb, redisJsonDb, sqlDb))}

	go messaging.DeleteCourseExams(repositories.NewExamRepositoryDb(redisDb, redisJsonDb))

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
	router.GET("/api/exam/get-exam/:examId/:userId", examHandler.GetExam)
	router.GET("/api/exam/is-student-take-exam-before/:examId/:userId", examHandler.IsStudentTakeExamBefore)
	router.DELETE("/api/exam/delete-exam/:examId", examHandler.DelExam)
	router.PUT("/api/exam/update-exam-info/:examId", examHandler.UpdateExamInfo)

	//question endpoints
	router.POST("/api/exam/add-question/:examId", questionHandler.Add)
	router.DELETE("/api/exam/delete-question/:examId/:questionId", questionHandler.Delete)
	router.PUT("/api/exam/update-question/:examId/:questionId", questionHandler.Update)

	//submit exam endpoint
	router.POST("/api/exam/submit", examSubmissionHandler.SubmitExam)

	//student grade endpoints
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
