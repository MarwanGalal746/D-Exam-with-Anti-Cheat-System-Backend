package repositories

import (
	"encoding/json"
	"exam_service/pkg/domain/models"
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"log"
	"strconv"
	"strings"
)

type QuestionRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
}

func (q QuestionRepositoryDb) Add(examId string, newQuestion models.Question) (*models.Question, error) {
	var examData models.ExamInfo
	key, err := q.redisJsonDb.JSONGet(examId, ".")
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
	qsArrLen, err := q.redisJsonDb.JSONArrLen(examData.ExamId, "questionIds")
	newQuestion.Id = examData.CourseId + "-" + examData.ExamId + "-" + strconv.Itoa(int(qsArrLen.(int64)))
	_, err = q.redisJsonDb.JSONSet(newQuestion.Id, ".", newQuestion)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrDb
	}
	_, err = q.redisJsonDb.JSONArrAppend(examData.ExamId, "questionIds", newQuestion.Id)
	if err != nil {
		log.Println(err)
		return nil, errs.ErrDb
	}
	return &newQuestion, nil
}

func NewQuestionRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler) QuestionRepositoryDb {
	return QuestionRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb}
}
