package exam

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type ExamRepositoryDb struct {
	db *redis.Client
}

var cursor uint64

func (e ExamRepositoryDb) Create(newExam Exam) error {
	examJson, err := json.Marshal(newExam)
	if err != nil {
		return err
	}
	return e.db.Set(newExam.Name, examJson, 0).Err()
}

func (e ExamRepositoryDb) Read() ([]Exam, error) {
	key := e.db.Scan(cursor, "*", 0).Iterator()
	allExams := make([]Exam, 0)
	for key.Next() {
		value := e.db.Get(key.Val())
		if err := value.Err(); err != nil {
			return nil, err
		}
		var exam Exam
		err := json.Unmarshal([]byte(value.Val()), &exam)
		if err != nil {
			return nil, err
		}
		allExams = append(allExams, exam)
	}
	if err := key.Err(); err != nil {
		return nil, err
	}
	return allExams, nil
}

func NewExamRepositoryDb(db *redis.Client) ExamRepositoryDb {
	return ExamRepositoryDb{db}
}
