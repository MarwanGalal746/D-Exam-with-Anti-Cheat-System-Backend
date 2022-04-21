package exam

import (
	"encoding/json"
	"exam_service/pkg/errs"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strings"
)

type ExamRepositoryDb struct {
	db *redis.Client
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	examJson, err := json.Marshal(newExam)
	if err != nil {
		return errs.ErrMarshallingInstance
	}

	// checking if there is an exam with the same name
	key := e.db.Scan(cursor, "*", 0).Iterator()
	for key.Next() {
		if key.Val() == newExam.Name {
			log.Println(err)
			return errs.ErrDuplicateExam
		}
	}
	if err := key.Err(); err != nil {
		log.Println(err)
		return errs.ErrDb
	}

	err = e.db.Set(newExam.Name, examJson, 0).Err()
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	return nil
}

func (e ExamRepositoryDb) GetAll() ([]Exam, error) {
	key := e.db.Scan(cursor, "*", 0).Iterator()
	allExams := make([]Exam, 0)

	for key.Next() {
		value := e.db.Get(key.Val())
		if err := value.Err(); err != nil {
			return nil, errs.ErrDb
		}
		var exam Exam
		err := json.Unmarshal([]byte(value.Val()), &exam)
		if err != nil {
			log.Println(err)
			return nil, errs.ErrUnmarshallingJson
		}
		allExams = append(allExams, exam)
	}
	if err := key.Err(); err != nil {
		log.Println(err)
		return nil, errs.ErrDb
	}

	return allExams, nil
}

func (e ExamRepositoryDb) GetExam(name string) (*Exam, error) {
	jsonExam := e.db.Get(name)
	fmt.Println(jsonExam.String())
	if strings.Contains(jsonExam.String(), "connect: connection refused") {
		return nil, errs.ErrDb
	} else if strings.Contains(jsonExam.String(), "redis: nil") {
		return nil, errs.ErrExamDoesNotExist
	}
	var exam *Exam
	err := json.Unmarshal([]byte(jsonExam.Val()), &exam)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrUnmarshallingJson
	}

	return exam, nil
}

func NewExamRepositoryDb(db *redis.Client) ExamRepositoryDb {
	return ExamRepositoryDb{db}
}
