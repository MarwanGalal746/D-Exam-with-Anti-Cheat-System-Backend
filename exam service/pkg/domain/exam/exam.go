package exam

import (
	"time"
)

type ExamInfo struct {
	ExamId      string    `json:"examId"`
	Name        string    `json:"name"`
	CourseId    string    `json:"courseId"`
	Duration    int       `json:"duration"`
	TotalMark   int       `json:"totalMark"`
	Date        time.Time `json:"date"`
	QuestionIds []string  `json:"questionIds,omitempty"`
}

type Exam struct {
	ExamData  ExamInfo   `json:"examData"`
	Questions []Question `json:"questions,omitempty"`
}

type ExamRepository interface {
	Create(Exam) error
	GetCourseExams(string) (*Course, error)
	GetExam(string) (*Exam, error)
	DelExam(string) error
}
