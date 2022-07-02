package models

type StudentGrade struct {
	Id             int    `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;unique;not null;type:int"`
	UserId         string `json:"userId,omitempty" gorm:"index:idx_submission,unique;not null;type:string"`
	ExamId         string `json:"examId,omitempty" gorm:"index:idx_submission,unique;not null;type:string"`
	CourseId       string `json:"courseId,omitempty" gorm:"<-;not null;type:string"`
	Grade          string `json:"grade" gorm:"<-;not null;type:string"`
	CheatingStatus string `json:"cheatingStatus" gorm:"<-;not null;type:string"`
}

type StudentGradeRepository interface {
	Add(string, string, string, Report) error
	GetAllStudentGrades(string) ([]Report, error)
	GetAllCourseGrades(string) ([]Report, error)
	GetAllExamGrades(string) ([]Report, error)
	GetUserCourseExamGrade(string, string, string) (*Report, error)
	DeleteAllStudentGrades(string) error
	DeleteAllCourseGrades(string) error
	DeleteAllExamGrades(string) error
	DeleteUserCourseExamGrade(string, string, string) error
}
