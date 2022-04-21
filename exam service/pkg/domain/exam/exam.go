package exam

type Exam struct {
	Name string `json:"name"`
}

type ExamRepository interface {
	Create(question Exam) error
	GetAll() ([]Exam, error)
	GetExam(string) (*Exam, error)
}
