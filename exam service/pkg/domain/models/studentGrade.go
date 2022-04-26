package models

type StudentGrade struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique;not null;type:int"`
	UserId         string `json:"userId" gorm:"index:idx_submission,unique;not null;type:string"`
	ExamId         string `json:"examId" gorm:"index:idx_submission,unique;not null;type:string"`
	CourseId       string `json:"courseId" gorm:"<-;not null;type:string"`
	Grade          string `json:"grade" gorm:"<-;not null;type:string"`
	CheatingStatus string `json:"cheatingStatus" gorm:"<-;not null;type:string"`
}
