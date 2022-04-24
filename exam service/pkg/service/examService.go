package service

import (
	"exam_service/pkg/domain/models"
)

type ExamService interface {
	Create(models.Exam) error
	GetCourseExams(string) (*models.Course, error)
	GetExam(string) (*models.Exam, error)
	DelExam(string) error
	UpdateExamInfo(string, models.ExamInfo) error
}

type DefaultExamService struct {
	repo models.ExamRepository
}

func (e DefaultExamService) Create(newExam models.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetCourseExams(courseId string) (*models.Course, error) {
	return e.repo.GetCourseExams(courseId)
}

func (e DefaultExamService) GetExam(examId string) (*models.Exam, error) {
	return e.repo.GetExam(examId)
}

func (e DefaultExamService) DelExam(examId string) error {
	return e.repo.DelExam(examId)
}

func (e DefaultExamService) UpdateExamInfo(examId string, newExam models.ExamInfo) error {
	return e.repo.UpdateExamInfo(examId, newExam)
}

func NewExamService(repository models.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
