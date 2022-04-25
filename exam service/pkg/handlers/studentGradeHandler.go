package handlers

import (
	"exam_service/pkg/service"
	"github.com/gin-gonic/gin"
)

type StudentGradeHandlers struct {
	service service.StudentGradeService
}

func (studentGradeHandler StudentGradeHandlers) Add(c *gin.Context) {
}
