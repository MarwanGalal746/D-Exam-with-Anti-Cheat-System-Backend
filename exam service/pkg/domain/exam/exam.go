package exam

type Exam struct {
	Name string `json:"name"`
}

type ExamRepository interface {
	Create(question Exam) error
	Read() ([]Exam, error)
}
