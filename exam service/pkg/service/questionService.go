package service

import "exam_service/pkg/domain/models"

type QuestionService interface {
	Add(string, models.Question) (*models.Question, error)
	Delete(string, string) error
	Update(string, string, models.Question) error
}

type DefaultQuestionService struct {
	repo models.QuestionRepository
}

func (q DefaultQuestionService) Add(examId string, newQuestion models.Question) (*models.Question, error) {
	return q.repo.Add(examId, newQuestion)
}

func (q DefaultQuestionService) Delete(examId string, questionId string) error {
	return q.repo.Delete(examId, questionId)
}

func (q DefaultQuestionService) Update(examId, questionId string, newQuestion models.Question) error {
	return q.repo.Update(examId, questionId, newQuestion)
}

func NewQuestionService(repository models.QuestionRepository) DefaultQuestionService {
	return DefaultQuestionService{repo: repository}
}
