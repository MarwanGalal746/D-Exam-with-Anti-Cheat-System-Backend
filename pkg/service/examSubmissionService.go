package service

import "exam_service/pkg/domain/models"

type ExamSubmissionService interface {
	SubmitExam(models.ExamSubmission) (*models.StudentScore, error)
}

type DefaultExamSubmissionService struct {
	repo models.ExamSubmissionRepository
}

func (e DefaultExamSubmissionService) SubmitExam(studentAns models.ExamSubmission) (*models.StudentScore, error) {
	return e.repo.SubmitExam(studentAns)
}

func NewExamSubmissionService(repository models.ExamSubmissionRepository) DefaultExamSubmissionService {
	return DefaultExamSubmissionService{repo: repository}
}
