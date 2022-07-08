package service

import (
	"exam_service/pkg/domain/models"
)

type ExamService interface {
	Create(models.Exam) error
	GetCourseExams([]string) ([]models.CourseExams, error)
	GetExam(string, string) (*models.Exam, error)
	DelExam(string) error
	DelCourseExams(string) error
	UpdateExamInfo(string, models.ExamInfo) error
}

type DefaultExamService struct {
	repo models.ExamRepository
}

func (e DefaultExamService) Create(newExam models.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetCourseExams(courseIds []string) ([]models.CourseExams, error) {
	return e.repo.GetCourseExams(courseIds)
}

func (e DefaultExamService) GetExam(examId string, userId string) (*models.Exam, error) {
	return e.repo.GetExam(examId, userId)
}

func (e DefaultExamService) DelExam(examId string) error {
	return e.repo.DelExam(examId)
}

func (e DefaultExamService) DelCourseExams(courseId string) error {
	return e.repo.DelCourseExams(courseId)
}

func (e DefaultExamService) UpdateExamInfo(examId string, newExam models.ExamInfo) error {
	return e.repo.UpdateExamInfo(examId, newExam)
}

func NewExamService(repository models.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
