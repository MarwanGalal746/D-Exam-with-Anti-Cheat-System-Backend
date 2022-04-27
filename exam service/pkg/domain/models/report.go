package models

type Report struct {
	Id              int          `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;unique;type:int"`
	Report          string       `json:"report" gorm:"<-;not null;type:text"`
	StudentGradeId  int          `json:"studentGradeId,omitempty" gorm:"<-;type:int"`
	StudentGradeObj StudentGrade `json:"studentGrade" gorm:"foreignKey:StudentGradeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
