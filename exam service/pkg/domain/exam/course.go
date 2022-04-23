package exam

type Course struct {
	CourseId string   `json:"courseId"`
	Exams    []string `json:"exams"`
}

type CourseDb struct {
	CourseId string `json:"courseId"`
	Exams    []Exam `json:"exams"`
}
