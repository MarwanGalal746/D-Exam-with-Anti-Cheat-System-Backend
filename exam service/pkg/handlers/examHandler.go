package handlers

import (
	"encoding/json"
	"exam_service/pkg/domain/exam"
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
	log.Println("Method: POST   Route: /api/exam/create-exam Function: Method: exam_service/pkg/handlers.ExamHandlers.Create")
	c.Writer.Header().Add("Content-Type", "application/json")
	var newExam exam.Exam
	_ = json.NewDecoder(c.Request.Body).Decode(&newExam)
	err := examHandler.service.Create(newExam)

	//handling errors
	if err != nil && err.Error() == errs.ErrMarshallingInstance.Error() {
		log.Println(errs.ErrMarshallingInstance.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrMarshallingInstance.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrDb.Error() {
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

func (examHandler ExamHandlers) GetAll(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	allExams, err := examHandler.service.GetAll()
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
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(allExams)
}

func (examHandler ExamHandlers) GetExam(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	exam, err := examHandler.service.GetExam(c.Param("name"))
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
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(exam)
}
