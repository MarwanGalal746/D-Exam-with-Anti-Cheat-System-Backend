package exam

import "time"

type Exam struct {
	ExamId    string        `json:"examId"`
	Name      string        `json:"name"`
	CourseId  string        `json:"courseId"`
	Date      time.Time     `json:"date"`
	Duration  time.Duration `json:"duration"`
	TotalMark int           `json:"totalMark"`
}

type ExamRepository interface {
	Create(Exam) error
	GetCourseExams(string) (*Course, error)
	//GetExam(string) (*Exam, error)
}
