package models

type CourseInfo struct {
	CourseId string   `json:"courseId" validate:"required"`
	ExamsIds []string `json:"examsIds" validate:"required"`
}

type Course struct {
	CourseData CourseInfo `json:"courseData" validate:"required"`
	ExamsData  []Exam     `json:"exams" validate:"required"`
}
