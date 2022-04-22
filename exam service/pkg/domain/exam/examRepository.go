package exam

import (
	"encoding/json"
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"log"
)

type ExamRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
}

type Course struct {
	CourseId string   `json:"courseId"`
	Exams    []string `json:"exams"`
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	// checking if there is an exam with the same name
	_, err := e.redisJsonDb.JSONGet(newExam.ExamId, ".")
	if err == nil {
		log.Println(err)
		return errs.ErrDuplicateExam
	}

	//check if this exam is not the first exam in the course
	ok, err := e.redisJsonDb.JSONGet(newExam.CourseId, ".")
	course := &Course{newExam.CourseId, []string{newExam.ExamId}}
	if ok == nil {
		ok, err = e.redisJsonDb.JSONSet(newExam.CourseId, ".", course)
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}
	} else {
		_, err := e.redisJsonDb.JSONArrAppend(newExam.CourseId, "exams", newExam.ExamId)
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}
	}

	_, err = e.redisJsonDb.JSONSet(newExam.ExamId, ".", newExam)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	return nil
}

func (e ExamRepositoryDb) GetCourseExams(courseId string) (*Course, error) {
	key := e.redisDb.Get(courseId)
	if err := key.Err(); err != nil {
		return nil, errs.ErrDb
	}
	var course Course
	err := json.Unmarshal([]byte(key.Val()), &course)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrUnmarshallingJson
	}
	return &course, nil
}

//
//func (e ExamRepositoryDb) GetExam(name string) (*Exam, error) {
//	jsonExam := e.db.Get(name)
//	if strings.Contains(jsonExam.String(), "connect: connection refused") {
//		return nil, errs.ErrDb
//	} else if strings.Contains(jsonExam.String(), "redis: nil") {
//		return nil, errs.ErrExamDoesNotExist
//	}
//	var exam *Exam
//	err := json.Unmarshal([]byte(jsonExam.Val()), &exam)
//	if err != nil {
//		log.Println(err)
//		return nil, errs.ErrUnmarshallingJson
//	}
//
//	return exam, nil
//}

func NewExamRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler) ExamRepositoryDb {
	return ExamRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb}
}
