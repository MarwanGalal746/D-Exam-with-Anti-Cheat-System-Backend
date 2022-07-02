package service

import (
	models2 "exam_service/domain/models"
)

type ExamService interface {
	Create(models2.Exam) error
	GetCourseExams([]string) ([]models2.CourseExams, error)
	GetExam(string) (*models2.Exam, error)
	DelExam(string) error
	DelCourseExams(string) error
	UpdateExamInfo(string, models2.ExamInfo) error
}

type DefaultExamService struct {
	repo models2.ExamRepository
}

func (e DefaultExamService) Create(newExam models2.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetCourseExams(courseIds []string) ([]models2.CourseExams, error) {
	return e.repo.GetCourseExams(courseIds)
}

func (e DefaultExamService) GetExam(examId string) (*models2.Exam, error) {
	return e.repo.GetExam(examId)
}

func (e DefaultExamService) DelExam(examId string) error {
	return e.repo.DelExam(examId)
}

func (e DefaultExamService) DelCourseExams(courseId string) error {
	return e.repo.DelCourseExams(courseId)
}

func (e DefaultExamService) UpdateExamInfo(examId string, newExam models2.ExamInfo) error {
	return e.repo.UpdateExamInfo(examId, newExam)
}

func NewExamService(repository models2.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
