package service

import "exam_service/pkg/domain/models"

type StudentGradeService interface {
	Add(string, string, string, models.StudentInfo) error
}

type DefaultStudentGradeService struct {
	repo models.StudentGradeRepository
}

func (s DefaultStudentGradeService) Add(userId, examId, courseId string, studentInfo models.StudentInfo) error {
	return s.repo.Add(userId, examId, courseId, studentInfo)
}

func NewStudentGradeService(repository models.StudentGradeRepository) DefaultStudentGradeService {
	return DefaultStudentGradeService{repo: repository}
}
