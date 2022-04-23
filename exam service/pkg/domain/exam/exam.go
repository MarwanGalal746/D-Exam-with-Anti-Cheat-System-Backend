package exam

import (
	"time"
)

type Exam struct {
	ExamId    string    `json:"examId"`
	Name      string    `json:"name"`
	CourseId  string    `json:"courseId"`
	Duration  int       `json:"duration"`
	TotalMark int       `json:"totalMark"`
	Date      time.Time `json:"date"`
}

type ExamRepository interface {
	Create(Exam) error
	GetCourseExams(string) (*CourseDb, error)
	//GetExam(string) (*Exam, error)
}

type Course struct {
	CourseId string   `json:"courseId"`
	Exams    []string `json:"exams"`
}

type CourseDb struct {
	CourseId string `json:"courseId"`
	Exams    []Exam `json:"exams"`
}
