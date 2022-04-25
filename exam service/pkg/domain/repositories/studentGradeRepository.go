package repositories

import (
	"gorm.io/gorm"
)

type StudentGradeRepositoryDb struct {
	sqlDb *gorm.DB
}

func (s StudentGradeRepositoryDb) Add() {

}

func NewStudentGradeRepositoryDb(sqlDb *gorm.DB) StudentGradeRepositoryDb {
	return StudentGradeRepositoryDb{sqlDb: sqlDb}
}
