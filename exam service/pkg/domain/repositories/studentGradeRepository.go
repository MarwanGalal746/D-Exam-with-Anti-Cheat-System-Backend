package repositories

import (
	"exam_service/pkg/domain/models"
	"exam_service/pkg/errs"
	"gorm.io/gorm"
	"log"
	"strings"
)

type StudentGradeRepositoryDb struct {
	sqlDb *gorm.DB
}

func (s StudentGradeRepositoryDb) Add(userId, examId, courseId string, studentInfo models.StudentInfo) error {
	student := &models.StudentGrade{
		UserId:         userId,
		CourseId:       courseId,
		ExamId:         examId,
		Grade:          studentInfo.Grade,
		CheatingStatus: studentInfo.CheatingStatus,
	}
	row := s.sqlDb.Create(&student)
	if row.Error != nil {
		log.Println(row.Error)
		if strings.Contains(row.Error.Error(), "duplicate key value") {
			return errs.ErrDuplicateUserExam
		}
		return errs.ErrDb
	}
	report := &models.Report{
		Report:          studentInfo.Report,
		StudentGradeId:  student.Id,
		StudentGradeObj: *student,
	}
	row = s.sqlDb.Create(report)
	if row.Error != nil {
		log.Println(row.Error)
		return errs.ErrDb
	}
	return nil
}

func NewStudentGradeRepositoryDb(sqlDb *gorm.DB) StudentGradeRepositoryDb {
	return StudentGradeRepositoryDb{sqlDb: sqlDb}
}
