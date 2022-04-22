package exam

type Exam struct {
	ExamId   string `json:"examId"`
	Name     string `json:"name"`
	CourseId string `json:"courseId"`
}

type ExamRepository interface {
	Create(question Exam) error
	//GetAll() ([]Exam, error)
	//GetExam(string) (*Exam, error)
}
