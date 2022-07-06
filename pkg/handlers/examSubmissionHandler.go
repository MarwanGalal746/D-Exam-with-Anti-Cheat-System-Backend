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

type ExamSubmissionHandlers struct {
	service service.ExamSubmissionService
}

func (examSubmissionHandler ExamSubmissionHandlers) SubmitExam(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var examSubmission models.ExamSubmission
	_ = json.NewDecoder(c.Request.Body).Decode(&examSubmission)
	err := validate.Struct(examSubmission)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}

	studentScore, err := examSubmissionHandler.service.SubmitExam(examSubmission)
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
	json.NewEncoder(c.Writer).Encode(studentScore)
}
