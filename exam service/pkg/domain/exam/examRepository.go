package exam

import (
	"exam_service/pkg/errs"
	"github.com/nitishm/go-rejson"
)

type ExamRepositoryDb struct {
	db *rejson.Handler
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	// checking if there is an exam with the same name
	_, err := e.db.JSONGet(newExam.Name, ".")
	if err == nil {
		return errs.ErrDuplicateExam
	}

	_, err = e.db.JSONSet(newExam.Name, ".", newExam)
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

func NewExamRepositoryDb(db *rejson.Handler) ExamRepositoryDb {
	return ExamRepositoryDb{db}
}

//you
//go redis
//db
