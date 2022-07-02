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

type QuestionHandlers struct {
	service service.QuestionService
}

func (questionHandler QuestionHandlers) Add(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var newQuestion models.Question
	_ = json.NewDecoder(c.Request.Body).Decode(&newQuestion)
	err := validate.Struct(newQuestion)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}
	questionAdded, err := questionHandler.service.Add(c.Param("examId"), newQuestion)

	//handling errors
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
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
	json.NewEncoder(c.Writer).Encode(questionAdded)
}

func (questionHandler QuestionHandlers) Delete(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")

	err := questionHandler.service.Delete(c.Param("examId"), c.Param("questionId"))
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrQuestionDoesNotExist.Error() {
		log.Println(errs.ErrQuestionDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrQuestionDoesNotExist.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("Question has been deleted successfully", http.StatusOK))
}

func (questionHandler QuestionHandlers) Update(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var newQuestion models.Question
	_ = json.NewDecoder(c.Request.Body).Decode(&newQuestion)
	err := validate.Struct(newQuestion)
	if err != nil {
		log.Println(errs.ErrRequiredFieldsAreMissed.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrRequiredFieldsAreMissed.Error(), http.StatusBadRequest))
		return
	}
	err = questionHandler.service.Update(c.Param("examId"), c.Param("questionId"), newQuestion)
	if err != nil && err.Error() == errs.ErrDb.Error() {
		log.Println(errs.ErrDb.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	} else if err != nil && err.Error() == errs.ErrQuestionDoesNotExist.Error() {
		log.Println(errs.ErrQuestionDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrQuestionDoesNotExist.Error(), http.StatusBadRequest))
		return
	} else if err != nil && err.Error() == errs.ErrExamDoesNotExist.Error() {
		log.Println(errs.ErrExamDoesNotExist.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrExamDoesNotExist.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(newQuestion)
}
