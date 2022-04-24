package service

import (
	"exam_service/pkg/domain/exam"
)

type ExamService interface {
	Create(exam.Exam) error
	GetCourseExams(string) (*exam.Course, error)
	GetExam(string) (*exam.Exam, error)
	DelExam(string) error
}

type DefaultExamService struct {
	repo exam.ExamRepository
}

func (e DefaultExamService) Create(newExam exam.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetCourseExams(courseId string) (*exam.Course, error) {
	return e.repo.GetCourseExams(courseId)
}

func (e DefaultExamService) GetExam(examId string) (*exam.Exam, error) {
	return e.repo.GetExam(examId)
}

func (e DefaultExamService) DelExam(examId string) error {
	return e.repo.DelExam(examId)
}

func NewExamService(repository exam.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
