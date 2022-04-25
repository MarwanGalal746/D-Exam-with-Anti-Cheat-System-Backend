package models

type StudentGrade struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique;not null;type:int"`
	UserId         string `json:"userId" gorm:"index:idx_submission,unique;not null;type:int"`
	ExamId         string `json:"examId" gorm:"index:idx_submission,unique;not null;type:int"`
	CourseId       string `json:"courseId" gorm:"<-;not null;type:int"`
	Grade          int    `json:"grade" gorm:"<-;not null;type:int"`
	CheatingStatus string `json:"cheatingStatus" gorm:"<-;not null;type:string"`
}

type StudentGradeRepository interface {
	Add()
}
