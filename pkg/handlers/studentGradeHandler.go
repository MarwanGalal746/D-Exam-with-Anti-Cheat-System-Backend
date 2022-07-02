package handlers

import (
	"encoding/json"
	"exam_service/pkg/domain/models"
	"exam_service/pkg/errs"
	"exam_service/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StudentGradeHandlers struct {
	service service.StudentGradeService
}

func (studentGradeHandler StudentGradeHandlers) Add(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var studentInfo models.Report
	_ = json.NewDecoder(c.Request.Body).Decode(&studentInfo)
	err := validate.Struct(studentInfo)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}

	err = studentGradeHandler.service.Add(c.Param("userId"),
		c.Param("examId"), c.Param("courseId"), studentInfo)
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The student grade has been added successfully", http.StatusOK))
}

func (studentGradeHandler StudentGradeHandlers) GetAllStudentGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	reports, err := studentGradeHandler.service.GetAllStudentGrades(c.Param("userId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(reports)
}

func (studentGradeHandler StudentGradeHandlers) GetAllCourseGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	reports, err := studentGradeHandler.service.GetAllCourseGrades(c.Param("courseId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(reports)
}

func (studentGradeHandler StudentGradeHandlers) GetAllExamGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	reports, err := studentGradeHandler.service.GetAllExamGrades(c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(reports)
}

func (studentGradeHandler StudentGradeHandlers) GetUserCourseExamGrade(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	report, err := studentGradeHandler.service.GetUserCourseExamGrade(c.Param("userId"),
		c.Param("courseId"), c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(report)
}

func (studentGradeHandler StudentGradeHandlers) DeleteAllStudentGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	err := studentGradeHandler.service.DeleteAllStudentGrades(c.Param("userId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The student grades have been deleted successfully", http.StatusOK))
}

func (studentGradeHandler StudentGradeHandlers) DeleteAllCourseGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	err := studentGradeHandler.service.DeleteAllCourseGrades(c.Param("courseId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The course grades have been deleted successfully", http.StatusOK))
}

func (studentGradeHandler StudentGradeHandlers) DeleteAllExamGrades(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	err := studentGradeHandler.service.DeleteAllExamGrades(c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The exam grades have been deleted successfully", http.StatusOK))
}

func (studentGradeHandler StudentGradeHandlers) DeleteUserCourseExamGrade(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	err := studentGradeHandler.service.DeleteUserCourseExamGrade(c.Param("userId"),
		c.Param("courseId"), c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The student grade has been deleted successfully", http.StatusOK))
}
