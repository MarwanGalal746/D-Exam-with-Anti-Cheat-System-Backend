package models

type Report struct {
	Id              int          `json:"id" gorm:"primaryKey;autoIncrement:true;unique;not null;type:uuid"`
	Report          string       `json:"report" gorm:"<-;not null;type:text"`
	StudentGradeId  int          `json:"studentGradeId" gorm:"<-;type:int"`
	StudentGradeObj StudentGrade `json:"studentGradeObj" gorm:"foreignKey:StudentGradeId"`
}
