package service

import (
	"exam_service/pkg/domain/exam"
)

type ExamService interface {
	Create(exam.Exam) error
	GetCourseExams(string) (*exam.CourseDb, error)
	//GetExam(string) (*exam.Exam, error)
}

type DefaultExamService struct {
	repo exam.ExamRepository
}

func (e DefaultExamService) Create(newExam exam.Exam) error {
	return e.repo.Create(newExam)
}

func (e DefaultExamService) GetCourseExams(courseId string) (*exam.CourseDb, error) {
	return e.repo.GetCourseExams(courseId)
}

//func (e DefaultExamService) GetExam(name string) (*exam.Exam, error) {
//	return e.repo.GetExam(name)
//}

func NewExamService(repository exam.ExamRepository) DefaultExamService {
	return DefaultExamService{repo: repository}
}
