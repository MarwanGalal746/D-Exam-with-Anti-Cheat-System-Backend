package exam

import (
	"encoding/json"
	"fmt"
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
	return e.db.Set(newExam.name, examJson, 0).Err()
}

func (e ExamRepositoryDb) Read() error {
	for {
		exams, cursor, err := e.db.Scan(cursor, "prefix:*", 0).Result()
		if err != nil {
			return err
		}
		for _, key := range exams {
			fmt.Println("key", key)
		}
		if cursor == 0 { // no more keys
			return nil
		}
	}
}
