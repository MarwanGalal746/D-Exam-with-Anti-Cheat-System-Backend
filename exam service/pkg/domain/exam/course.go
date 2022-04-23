package exam

type CourseInfo struct {
	CourseId string   `json:"courseId"`
	ExamsIds []string `json:"examsIds"`
}

type Course struct {
	CourseData CourseInfo `json:"courseData"`
	ExamsData  []ExamInfo `json:"exams"`
}
