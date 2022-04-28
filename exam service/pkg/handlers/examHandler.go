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

type ExamHandlers struct {
	service service.ExamService
}

func (examHandler ExamHandlers) Create(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var newExam models.Exam
	_ = json.NewDecoder(c.Request.Body).Decode(&newExam)
	err := validate.Struct(newExam)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}
	err = examHandler.service.Create(newExam)

	//handling errors
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrDuplicateExam.Error() {
		log.Println(errs.ErrDuplicateExam.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDuplicateExam.Error(), http.StatusBadRequest))
		return
	}

	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(newExam)
}

func (examHandler ExamHandlers) GetCourseExams(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	allExams, err := examHandler.service.GetCourseExams(c.Param("courseId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrUnmarshallingJson.Error() {
		log.Println(errs.ErrUnmarshallingJson.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrUnmarshallingJson.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrCourseDoesNotExist.Error() {
		log.Println(errs.ErrCourseDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrCourseDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(allExams)
}

func (examHandler ExamHandlers) GetExam(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	allExams, err := examHandler.service.GetExam(c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrUnmarshallingJson.Error() {
		log.Println(errs.ErrUnmarshallingJson.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrUnmarshallingJson.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(allExams)
}

func (examHandler ExamHandlers) DelExam(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	err := examHandler.service.DelExam(c.Param("examId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrUnmarshallingJson.Error() {
		log.Println(errs.ErrUnmarshallingJson.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrUnmarshallingJson.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("Exam has been deleted successfully", http.StatusOK))
}

func (examHandler ExamHandlers) UpdateExamInfo(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var newExam models.ExamInfo
	_ = json.NewDecoder(c.Request.Body).Decode(&newExam)
	err := validate.Struct(newExam)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}
	err = examHandler.service.UpdateExamInfo(c.Param("examId"), newExam)
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrUnmarshallingJson.Error() {
		log.Println(errs.ErrUnmarshallingJson.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrUnmarshallingJson.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamUpdateId.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamUpdateId.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(newExam)
}
