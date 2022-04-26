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

func (s StudentGradeRepositoryDb) Add(userId, examId, courseId string, studentInfo models.Report) error {
	studentInfo.StudentGradeObj.UserId = userId
	studentInfo.StudentGradeObj.CourseId = courseId
	studentInfo.StudentGradeObj.ExamId = examId
	row := s.sqlDb.Create(&studentInfo.StudentGradeObj)
	if row.Error != nil {
		log.Println(row.Error)
		if strings.Contains(row.Error.Error(), "duplicate key value") {
			return errs.ErrDuplicateUserExam
		}
		return errs.ErrDb
	}
	row = s.sqlDb.Create(&studentInfo)
	if row.Error != nil {
		log.Println(row.Error)
		return errs.ErrDb
	}
	return nil
}

func (s StudentGradeRepositoryDb) GetAllStudentGrades(userId string) ([]models.Report, error) {
	rows, err := s.sqlDb.Raw("Select * from student_grades join reports on student_grades.id = reports.student_grade_id where student_grades.user_id=?",
		userId).Rows()
	if err != nil {
		return nil, errs.ErrDb
	}
	var reports []models.Report
	for rows.Next() {
		var report models.Report
		err := rows.Scan(&report.StudentGradeId, &report.StudentGradeObj.UserId, &report.StudentGradeObj.ExamId,
			&report.StudentGradeObj.CourseId,
			&report.StudentGradeObj.Grade, &report.StudentGradeObj.CheatingStatus, &report.Id, &report.Report,
			&report.StudentGradeId)
		if err != nil {
			return nil, errs.ErrDb
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func NewStudentGradeRepositoryDb(sqlDb *gorm.DB) StudentGradeRepositoryDb {
	return StudentGradeRepositoryDb{sqlDb: sqlDb}
}
