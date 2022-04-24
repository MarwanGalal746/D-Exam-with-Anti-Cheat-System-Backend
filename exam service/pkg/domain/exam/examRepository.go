package exam

import (
	"encoding/json"
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"log"
	"strconv"
	"strings"
)

type ExamRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	// checking if there is an exam with the same name
	_, err := e.redisJsonDb.JSONGet(newExam.ExamData.ExamId, ".")
	if err == nil {
		log.Println(err)
		return errs.ErrDuplicateExam
	}

	//check if this exam is not the first exam in the course
	ok, err := e.redisJsonDb.JSONGet(newExam.ExamData.CourseId, ".")
	course := &CourseInfo{CourseId: newExam.ExamData.CourseId, ExamsIds: []string{newExam.ExamData.ExamId}}
	if ok == nil {
		ok, err = e.redisJsonDb.JSONSet(newExam.ExamData.CourseId, ".", course)
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}
	} else {
		_, err := e.redisJsonDb.JSONArrAppend(newExam.ExamData.CourseId, "examsIds", newExam.ExamData.ExamId)
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}
	}

	//setting questions Ids in examData and storing questions in database
	i := 0
	for _, qs := range newExam.Questions {
		qs.Id = newExam.ExamData.CourseId + "-" + newExam.ExamData.ExamId + "-" + strconv.Itoa(i)
		newExam.ExamData.QuestionIds = append(newExam.ExamData.QuestionIds, qs.Id)
		ok, err = e.redisJsonDb.JSONSet(qs.Id, ".", qs)
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}
		i++
	}
	_, err = e.redisJsonDb.JSONSet(newExam.ExamData.ExamId, ".", newExam.ExamData)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	return nil
}

func (e ExamRepositoryDb) GetCourseExams(courseId string) (*Course, error) {
	var course Course
	key, err := e.redisJsonDb.JSONGet(courseId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return nil, errs.ErrCourseDoesNotExist
		}
		return nil, errs.ErrDb
	}
	err = json.Unmarshal(key.([]byte), &course.CourseData)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrUnmarshallingJson
	}
	for _, examId := range course.CourseData.ExamsIds {
		key, err := e.redisJsonDb.JSONGet(examId, ".")
		if err != nil {
			log.Println(err)
			return nil, errs.ErrDb
		}
		var exam ExamInfo
		err = json.Unmarshal(key.([]byte), &exam)
		if err != nil {
			log.Println(err)
			return nil, errs.ErrUnmarshallingJson
		}
		//this line to make the array of questions id empty
		//because it's not important and secure to show questions id to the user in this endpoint
		exam.QuestionIds = []string{}
		course.ExamsData = append(course.ExamsData, exam)
	}
	return &course, nil
}

func (e ExamRepositoryDb) GetExam(examId string) (*Exam, error) {
	var examData ExamInfo
	key, err := e.redisJsonDb.JSONGet(examId, ".")
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
	var exam Exam
	exam.ExamData = examData
	for _, qsId := range examData.QuestionIds {
		key, err := e.redisJsonDb.JSONGet(qsId, ".")
		if err != nil {
			log.Println(err)
			return nil, errs.ErrDb
		}
		var qs Question
		err = json.Unmarshal(key.([]byte), &qs)
		if err != nil {
			log.Println(err)
			return nil, errs.ErrUnmarshallingJson
		}
		exam.Questions = append(exam.Questions, qs)
	}
	return &exam, nil
}

func NewExamRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler) ExamRepositoryDb {
	return ExamRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb}
}
