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

type ExamRepositoryDb struct {
	redisDb     *redis.Client
	redisJsonDb *rejson.Handler
}

func (e ExamRepositoryDb) Create(newExam models.Exam) error {
	// checking if there is an exam with the same name
	_, err := e.redisJsonDb.JSONGet(newExam.ExamData.ExamId, ".")
	if err == nil {
		log.Println(err)
		return errs.ErrDuplicateExam
	}

	//check if this exam is not the first exam in the course
	ok, err := e.redisJsonDb.JSONGet(newExam.ExamData.CourseId, ".")
	course := &models.CourseInfo{CourseId: newExam.ExamData.CourseId, ExamsIds: []string{newExam.ExamData.ExamId}}
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

func (e ExamRepositoryDb) GetCourseExams(courseId string) (*models.Course, error) {
	var course models.Course
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
		var exam models.ExamInfo
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

func (e ExamRepositoryDb) GetExam(examId string) (*models.Exam, error) {
	var examData models.ExamInfo
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
	return &exam, nil
}

func (e ExamRepositoryDb) DelExam(examId string) error {
	//get the exam information
	var examData models.ExamInfo
	key, err := e.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrExamDoesNotExist
		}
		return errs.ErrDb
	}
	err = json.Unmarshal(key.([]byte), &examData)
	if err != nil {
		log.Println(err)
		return errs.ErrUnmarshallingJson
	}

	//removing the questions of the exam first
	for _, qsId := range examData.QuestionIds {
		//delete the question itself
		key, err = e.redisJsonDb.JSONDel(qsId, ".")
		if err != nil {
			log.Println(err)
			return errs.ErrDb
		}

		// this block of code will be wanted in the future
		////get the index of the qs in examInfo
		//qsInd, err := e.redisJsonDb.JSONArrIndex(examData.ExamId, "questionIds", qsId)
		//if err != nil {
		//	log.Println(err)
		//	return errs.ErrDb
		//}
		//if err != nil {
		//	return err
		//}
		////var ind int64
		////ind = int
		////removing the qs id from exam info
		//key, err = e.redisJsonDb.JSONArrPop(examData.ExamId, "questionIds", int(qsInd.(int64)))
		//if err != nil {
		//	log.Println(err)
		//	return errs.ErrDb
		//}
	}

	//removing the exam from the course information
	//get the index of the course
	courseId, err := e.redisJsonDb.JSONArrIndex(examData.CourseId, "examsIds", examData.ExamId)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	//removing the exam id from course info
	key, err = e.redisJsonDb.JSONArrPop(examData.CourseId, "examsIds", int(courseId.(int64)))
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}

	//removing the exam itself
	key, err = e.redisJsonDb.JSONDel(examData.ExamId, ".")
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	return nil
}

func (e ExamRepositoryDb) UpdateExamInfo(examId string, newExam models.ExamInfo) error {
	var examData models.ExamInfo
	key, err := e.redisJsonDb.JSONGet(examId, ".")
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), errs.ErrRedisNil.Error()) {
			return errs.ErrExamDoesNotExist
		}
		return errs.ErrDb
	}
	err = json.Unmarshal(key.([]byte), &examData)
	if err != nil {
		log.Println(err)
		return errs.ErrUnmarshallingJson
	}
	if newExam.ExamId != examData.ExamId {
		return errs.ErrExamUpdateId
	}
	updatedExam := models.ResetExamInfo(newExam)
	updatedExam.QuestionIds = examData.QuestionIds
	_, err = e.redisJsonDb.JSONSet(updatedExam.ExamId, ".", updatedExam)
	if err != nil {
		log.Println(err)
		return errs.ErrDb
	}
	//this line to make the array of questions id empty
	//because it's not important and secure to show questions id to the user in this endpoint
	updatedExam.QuestionIds = []string{}

	return nil
}

func NewExamRepositoryDb(redisDb *redis.Client, redisJsonDb *rejson.Handler) ExamRepositoryDb {
	return ExamRepositoryDb{redisDb: redisDb, redisJsonDb: redisJsonDb}
}
