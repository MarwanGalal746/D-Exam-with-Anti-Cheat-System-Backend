package service

import "exam_service/pkg/domain/exam"

type ExamService interface {
	Create(exam.Exam) (int, error)
	Read() error
}

type DefaultExamService struct {
	repo exam.ExamRepository
}

func (e DefaultExamService) Create(newExam exam.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) Read() error {
	return e.repo.Read()
}

func NewExamService(repository exam.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
