package exam

import (
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

type ExamRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
}

type courseDb struct {
	CourseId string   `json:"courseId"`
	Exams    []string `json:"exams"`
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	// checking if there is an exam with the same name
	_, err := e.redisJsonDb.JSONGet(newExam.ExamId, ".")
	if err == nil {
		return errs.ErrDuplicateExam
	}

	//check if this exam is not the first exam in the course
	ok, err := e.redisJsonDb.JSONGet(newExam.CourseId, ".")
	course := &courseDb{newExam.CourseId, []string{newExam.ExamId}}
	if ok == nil {
		ok, err = e.redisJsonDb.JSONSet(newExam.CourseId, ".", course)
		if err != nil {
			return errs.ErrDb
		}
	} else {
		_, err := e.redisJsonDb.JSONArrAppend(newExam.CourseId, "exams", newExam.ExamId)
		if err != nil {
			return errs.ErrDb
		}
	}

	_, err = e.redisJsonDb.JSONSet(newExam.ExamId, ".", newExam)
	if err != nil {
		return errs.ErrDb
	}
	return nil
}

//func (e ExamRepositoryDb) GetAll() ([]Exam, error) {
//	key := e.db.Scan(cursor, "*", 0).Iterator()
//	allExams := make([]Exam, 0)
//
//	for key.Next() {
//		value := e.db.Get(key.Val())
//		if err := value.Err(); err != nil {
//			return nil, errs.ErrDb
//		}
//		var exam Exam
//		err := json.Unmarshal([]byte(value.Val()), &exam)
//		if err != nil {
//			log.Println(err)
//			return nil, errs.ErrUnmarshallingJson
//		}
//		allExams = append(allExams, exam)
//	}
//	if err := key.Err(); err != nil {
//		log.Println(err)
//		return nil, errs.ErrDb
//	}
//
//	return allExams, nil
//}

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
