package service

import "exam_service/pkg/domain/models"

type StudentGradeService interface {
	Add(string, string, string, models.Report) error
	GetAllStudentGrades(string) ([]models.Report, error)
	GetAllCourseGrades(string) ([]models.Report, error)
}

type DefaultStudentGradeService struct {
	repo models.StudentGradeRepository
}

func (s DefaultStudentGradeService) Add(userId, examId, courseId string, studentInfo models.Report) error {
	return s.repo.Add(userId, examId, courseId, studentInfo)
}

func (s DefaultStudentGradeService) GetAllStudentGrades(userId string) ([]models.Report, error) {
	return s.repo.GetAllStudentGrades(userId)
}

func (s DefaultStudentGradeService) GetAllCourseGrades(courseId string) ([]models.Report, error) {
	return s.repo.GetAllCourseGrades(courseId)
}

func NewStudentGradeService(repository models.StudentGradeRepository) DefaultStudentGradeService {
	return DefaultStudentGradeService{repo: repository}
}
