package exam

import "exam_service/pkg/domain/question"

type Exam struct {
	Name      string              `json:"name"`
	CourseId  int                 `json:"courseId"`
	Questions []question.Question `json:"questions"`
}

type ExamRepository interface {
	Create(question Exam) error
	//GetAll() ([]Exam, error)
	//GetExam(string) (*Exam, error)
}
