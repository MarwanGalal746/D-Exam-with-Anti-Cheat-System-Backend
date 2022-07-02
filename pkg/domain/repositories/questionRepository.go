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

func (q QuestionRepositoryDb) Delete(examId string, questionId string) error {
	//check if there is exam with that id
	_, err := q.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrExamDoesNotExist
		}
		return errs.ErrDb
	}

	//check if there is question with that id
	_, err = q.redisJsonDb.JSONGet(questionId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrQuestionDoesNotExist
		}
		return errs.ErrDb
	}

	// this block of code will be wanted in the future
	//get the index of the qs in examInfo
	qsInd, err := q.redisJsonDb.JSONArrIndex(examId, "questionIds", questionId)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	//removing the qs id from exam info
	_, err = q.redisJsonDb.JSONArrPop(examId, "questionIds", int(qsInd.(int64)))
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}

	//removing the question itself
	_, err = q.redisJsonDb.JSONDel(questionId, ".")
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}

	return nil
}

func (q QuestionRepositoryDb) Update(examId, questionId string, newQuestion models.Question) error {
	//check if there is exam with that id
	_, err := q.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrExamDoesNotExist
		}
		return errs.ErrDb
	}

	//check if there is question with that id
	_, err = q.redisJsonDb.JSONGet(questionId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrQuestionDoesNotExist
		}
		return errs.ErrDb
	}

	newQuestion.Id = questionId

	_, err = q.redisJsonDb.JSONSet(newQuestion.Id, ".", newQuestion)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}

	return nil
}

func NewQuestionRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler) QuestionRepositoryDb {
	return QuestionRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb}
}
