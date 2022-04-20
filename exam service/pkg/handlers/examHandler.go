package handlers

import (
	"encoding/json"
	"exam_service/pkg/domain/exam"
	"exam_service/pkg/errs"
	"exam_service/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExamHandlers struct {
	service service.ExamService
}

func (examHandler ExamHandlers) Create(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	var newExam exam.Exam
	fmt.Println(newExam.Id, " ", newExam.Name)
	_ = json.NewDecoder(c.Request.Body).Decode(&newExam)
	fmt.Println(newExam.Id, " ", newExam.Name)
	err := examHandler.service.Create(newExam)
	newExam.Id = int(uint(0))
	//handling errors
	if err != nil && err.Error() == `ERROR: duplicate key value violates unique constraint "sandwiches_name_key" (SQLSTATE 23505)` {
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDuplicateValue.Error(), http.StatusBadRequest))
		return
	} else if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(newExam)
}

func (examHandler ExamHandlers) Read(c *gin.Context) {
	c.Writer.Header().Add("Content-Type", "application/json")
	allExams, err := examHandler.service.Read()
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDb.Error(), http.StatusInternalServerError))
		return
	}
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(allExams)
}
