package exam

type CourseInfo struct {
	CourseId string   `json:"courseId"`
	ExamsIds []string `json:"examsIds"`
}

type Course struct {
	CourseDate CourseInfo `json:"courseDate"`
	Exams      []string   `json:"exams"`
}
