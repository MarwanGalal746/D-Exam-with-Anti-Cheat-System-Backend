package models

type CourseInfo struct {
	CourseId string   `json:"courseId" validate:"required"`
	ExamsIds []string `json:"examsIds" validate:"required"`
}

type Course struct {
	CourseData CourseInfo `json:"courseData" validate:"required"`
	ExamsData  []ExamInfo `json:"exams" validate:"required"`
}
