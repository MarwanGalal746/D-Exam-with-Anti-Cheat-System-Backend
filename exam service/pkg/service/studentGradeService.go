package service

import "exam_service/pkg/domain/models"

type StudentGradeService interface {
	Add(string, string, string, map[string]interface{}) error
}

type DefaultStudentGradeService struct {
	repo models.StudentGradeRepository
}

func (s DefaultStudentGradeService) Add(userId, examId, courseId string, data map[string]interface{}) error {
	return s.repo.Add(userId, examId, courseId, data)
}

func NewStudentGradeService(repository models.StudentGradeRepository) DefaultStudentGradeService {
	return DefaultStudentGradeService{repo: repository}
}
