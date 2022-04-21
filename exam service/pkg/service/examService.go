package service

import "exam_service/pkg/domain/exam"

type ExamService interface {
	Create(exam.Exam) error
	GetAll() ([]exam.Exam, error)
}

type DefaultExamService struct {
	repo exam.ExamRepository
}

func (e DefaultExamService) Create(newExam exam.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetAll() ([]exam.Exam, error) {
	return e.repo.GetAll()
}

func NewExamService(repository exam.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
