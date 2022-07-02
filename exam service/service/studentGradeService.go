package service

import (
	models2 "exam_service/domain/models"
)

type StudentGradeService interface {
	Add(string, string, string, models2.Report) error
	GetAllStudentGrades(string) ([]models2.Report, error)
	GetAllCourseGrades(string) ([]models2.Report, error)
	GetAllExamGrades(string) ([]models2.Report, error)
	GetUserCourseExamGrade(string, string, string) (*models2.Report, error)
	DeleteAllStudentGrades(string) error
	DeleteAllCourseGrades(string) error
	DeleteAllExamGrades(string) error
	DeleteUserCourseExamGrade(string, string, string) error
}

type DefaultStudentGradeService struct {
	repo models2.StudentGradeRepository
}

func (s DefaultStudentGradeService) Add(userId, examId, courseId string, studentInfo models2.Report) error {
	return s.repo.Add(userId, examId, courseId, studentInfo)
}

func (s DefaultStudentGradeService) GetAllStudentGrades(userId string) ([]models2.Report, error) {
	return s.repo.GetAllStudentGrades(userId)
}

func (s DefaultStudentGradeService) GetAllCourseGrades(courseId string) ([]models2.Report, error) {
	return s.repo.GetAllCourseGrades(courseId)
}

func (s DefaultStudentGradeService) GetAllExamGrades(examId string) ([]models2.Report, error) {
	return s.repo.GetAllExamGrades(examId)
}

func (s DefaultStudentGradeService) GetUserCourseExamGrade(userId, courseId, examId string) (*models2.Report, error) {
	return s.repo.GetUserCourseExamGrade(userId, courseId, examId)
}

func (s DefaultStudentGradeService) DeleteAllStudentGrades(userId string) error {
	return s.repo.DeleteAllStudentGrades(userId)
}

func (s DefaultStudentGradeService) DeleteAllCourseGrades(courseId string) error {
	return s.repo.DeleteAllCourseGrades(courseId)
}

func (s DefaultStudentGradeService) DeleteAllExamGrades(examId string) error {
	return s.repo.DeleteAllExamGrades(examId)
}

func (s DefaultStudentGradeService) DeleteUserCourseExamGrade(userId, courseId, examId string) error {
	return s.repo.DeleteUserCourseExamGrade(userId, courseId, examId)
}

func NewStudentGradeService(repository models2.StudentGradeRepository) DefaultStudentGradeService {
	return DefaultStudentGradeService{repo: repository}
}
