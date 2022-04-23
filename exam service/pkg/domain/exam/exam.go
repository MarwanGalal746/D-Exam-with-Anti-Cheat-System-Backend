package exam

import "time"

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
	GetCourseExams(string) (*Course, error)
	//GetExam(string) (*Exam, error)
}
