package exam

type Exam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ExamRepository interface {
	Create(question Exam) error
	Read() error
	//Update(Exam, string) error
	//Delete(string) error
}
