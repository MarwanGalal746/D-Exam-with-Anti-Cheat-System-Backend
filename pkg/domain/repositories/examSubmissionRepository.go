package repositories

import (
	"encoding/json"
	"exam_service/pkg/domain/models"
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"gorm.io/gorm"
	"log"
	"strings"
)

type ExamSubmissionRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
	sqlDb       *gorm.DB
}

func (e ExamSubmissionRepositoryDb) SubmitExam(studentAns models.ExamSubmission) (*models.StudentScore, error) {
	var examData models.ExamInfo
	key, err := e.redisJsonDb.JSONGet(studentAns.ExamId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return nil, errs.ErrExamDoesNotExist
		}
		return nil, errs.ErrDb
	}
	err = json.Unmarshal(key.([]byte), &examData)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrUnmarshallingJson
	}
	var exam models.Exam
	exam.ExamData = examData
	for _, qsId := range examData.QuestionIds {
		key, err := e.redisJsonDb.JSONGet(qsId, ".")
		if err != nil {
			log.Println(err)
			return nil, errs.ErrDb
		}
		var qs models.Question
		err = json.Unmarshal(key.([]byte), &qs)
		if err != nil {
			log.Println(err)
			return nil, errs.ErrUnmarshallingJson
		}
		exam.Questions = append(exam.Questions, qs)
	}
	var qsMark float64
	qsMark = float64(exam.ExamData.TotalMark / len(examData.QuestionIds))
	var studentMark float64
	studentMark = 0.0
	for i := 0; i < len(studentAns.Answers); i++ {
		for r := 0; r < len(exam.Questions); r++ {
			if studentAns.Answers[i].QsId == exam.Questions[i].Id {
				if studentAns.Answers[i].Answer == exam.Questions[i].RightChoice {
					studentMark += qsMark
				}
				break
			}
		}
	}
	studentInfo := &models.Report{
		Report: studentAns.Report,
		StudentGradeObj: models.StudentGrade{
			UserId:         studentAns.UserId,
			CourseId:       studentAns.CourseId,
			ExamId:         studentAns.ExamId,
			Grade:          studentMark,
			CheatingStatus: studentAns.CheatingStatus,
		},
	}
	row := e.sqlDb.Create(&studentInfo.StudentGradeObj)
	if row.Error != nil {
		log.Println(row.Error)
		if strings.Contains(row.Error.Error(), "duplicate key value") {
			return nil, errs.ErrDuplicateUserExam
		}
		return nil, errs.ErrDb
	}
	row = e.sqlDb.Create(&studentInfo)
	if row.Error != nil {
		log.Println(row.Error)
		return nil, errs.ErrDb
	}
	return &models.StudentScore{Score: studentMark, Total: exam.ExamData.TotalMark}, nil
}

func NewExamSubmissionRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler, sqlDb *gorm.DB) ExamSubmissionRepositoryDb {
	return ExamSubmissionRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb, sqlDb: sqlDb}
}
