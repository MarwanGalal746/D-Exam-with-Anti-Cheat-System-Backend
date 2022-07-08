package repositories

import (
	"exam_service/pkg/domain/models"
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"gorm.io/gorm"
	"log"
	"strings"
)

type StudentGradeRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
	sqlDb       *gorm.DB
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

func (s StudentGradeRepositoryDb) GetAllCourseGrades(courseId string) ([]models.Report, error) {
	rows, err := s.sqlDb.Raw("Select * from student_grades join reports on student_grades.id = reports.student_grade_id where student_grades.course_id=?",
		courseId).Rows()
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

func (s StudentGradeRepositoryDb) GetAllExamGrades(examId string) ([]models.Report, error) {

	// validate if exam exists or not
	_, err := s.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return nil, errs.ErrExamDoesNotExist
		}
		return nil, errs.ErrDb
	}

	rows, err := s.sqlDb.Raw("Select * from student_grades join reports on student_grades.id = reports.student_grade_id where student_grades.exam_id=?",
		examId).Rows()
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

func (s StudentGradeRepositoryDb) GetUserCourseExamGrade(userId, courseId, examId string) (*models.Report, error) {
	// validate if exam exists or not
	_, err := s.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return nil, errs.ErrExamDoesNotExist
		}
		return nil, errs.ErrDb
	}

	// validate if course exists or not
	_, err = s.redisJsonDb.JSONGet(courseId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return nil, errs.ErrCourseDoesNotExist
		}
		return nil, errs.ErrDb
	}

	row := s.sqlDb.Raw("Select * from student_grades join reports on student_grades.id = reports.student_grade_id where student_grades.user_id=? AND student_grades.course_id=? AND student_grades.exam_id=?",
		userId, courseId, examId).Row()
	if row.Err() != nil {
		log.Println(row.Err().Error())
		if strings.Contains(row.Err().Error(), "no rows in result set") {
			return nil, errs.ErrStudentDoesNotTakeTheExam
		} else {
			return nil, errs.ErrDb
		}
	}
	var report models.Report

	err = row.Scan(&report.StudentGradeId, &report.StudentGradeObj.UserId, &report.StudentGradeObj.ExamId,
		&report.StudentGradeObj.CourseId,
		&report.StudentGradeObj.Grade, &report.StudentGradeObj.CheatingStatus, &report.Id, &report.Report,
		&report.StudentGradeId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errs.ErrStudentDoesNotTakeTheExam
		} else {
			log.Println(err)
			return nil, errs.ErrDb
		}
	}
	return &report, nil
}

func (s StudentGradeRepositoryDb) DeleteAllStudentGrades(userId string) error {
	row := s.sqlDb.Raw(`DELETE FROM student_grades where student_grades.user_id=?`,
		userId).Row()
	if row.Err() != nil {
		log.Println(row.Err())
		return errs.ErrDb
	}
	return nil
}

func (s StudentGradeRepositoryDb) DeleteAllCourseGrades(courseId string) error {
	row := s.sqlDb.Raw(`DELETE FROM student_grades where student_grades.course_id=?`,
		courseId).Row()
	if row.Err() != nil {
		log.Println(row.Err())
		return errs.ErrDb
	}
	return nil
}

func (s StudentGradeRepositoryDb) DeleteAllExamGrades(examId string) error {
	row := s.sqlDb.Raw(`DELETE FROM student_grades where student_grades.exam_id=?`,
		examId).Row()
	if row.Err() != nil {
		log.Println(row.Err())
		return errs.ErrDb
	}
	return nil
}

func (s StudentGradeRepositoryDb) DeleteUserCourseExamGrade(userId, courseId, examId string) error {
	row := s.sqlDb.Raw(`DELETE FROM student_grades WHERE student_grades.user_id=? AND student_grades.course_id=? AND student_grades.exam_id=?`,
		userId, courseId, examId).Row()
	if row.Err() != nil {
		log.Println(row.Err())
		return errs.ErrDb
	}
	return nil
}

func NewStudentGradeRepositoryDb(sqlDb *gorm.DB, redisDb *redis.Client, redisJsonDb *rejson.Handler) StudentGradeRepositoryDb {
	return StudentGradeRepositoryDb{sqlDb: sqlDb, redisDb: redisDb, redisJsonDb: redisJsonDb}
}
