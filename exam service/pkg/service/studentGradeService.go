package service

import "exam_service/pkg/domain/models"

type StudentGradeService interface {
	Add()
}

type DefaultStudentGradeService struct {
	repo models.StudentGradeRepository
}

func (s DefaultStudentGradeService) Add() {
	s.repo.Add()
}

func NewStudentGradeService(repository models.StudentGradeRepository) DefaultStudentGradeService {
	return DefaultStudentGradeService{repo: repository}
}
