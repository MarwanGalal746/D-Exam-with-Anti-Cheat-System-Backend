package models

import (
	"time"
)

type Exam struct {
	ExamId    string     `json:"examId" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	CourseId  string     `json:"courseId" validate:"required"`
	Duration  int        `json:"duration" validate:"required"`
	TotalMark int        `json:"totalMark" validate:"required"`
	Date      time.Time  `json:"date" validate:"required"`
	Questions []Question `json:"questions,omitempty"`
}

func ResetExamInfo(source Exam) Exam {
	return Exam{ExamId: source.ExamId, Name: source.Name, CourseId: source.CourseId, Duration: source.Duration,
		TotalMark: source.TotalMark, Date: source.Date}
}

type ExamRepository interface {
	Create(Exam) error
	GetCourseExams(string) (*Course, error)
	GetExam(string) (*Exam, error)
	DelExam(string) error
	UpdateExamInfo(string, Exam) error
}
