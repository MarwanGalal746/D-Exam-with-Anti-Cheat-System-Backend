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
	var studentInfo models.StudentInfo
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
	} else if err != nil && err.Error() == errs.ErrDuplicateUserExam.Error() {
		log.Println(errs.ErrDuplicateUserExam.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(c.Writer).Encode(errs.NewResponse(errs.ErrDuplicateUserExam.Error(), http.StatusBadRequest))
		return
	}
	//sending the response
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.WriteHeader(http.StatusOK)
	//sending the response
	c.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(c.Writer).Encode(errs.NewResponse("The student grade has been added successfully", http.StatusOK))
}
