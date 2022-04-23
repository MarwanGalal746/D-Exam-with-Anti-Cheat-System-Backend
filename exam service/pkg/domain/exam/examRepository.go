package exam

import (
	"exam_service/pkg/errs"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
	"log"
	"strconv"
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

//func (e ExamRepositoryDb) GetCourseExams(courseId string) (*CourseDb, error) {
//	var courseDb CourseDb
//	key, err := e.redisJsonDb.JSONGet(courseId, ".")
//	if err != nil {
//		log.Println(err)
//		return nil, errs.ErrDb
//	}
//	var course Course
//	err = json.Unmarshal(key.([]byte), &course)
//	if err != nil {
//		log.Println(err)
//		return nil, errs.ErrUnmarshallingJson
//	}
//	courseDb.CourseId = course.CourseId
//	fmt.Println(len(courseDb.Exams))
//	//for _, examName := range course.Exams {
//	//	key, err = e.redisJsonDb.JSONGet(examName, ".")
//	//	if err != nil {
//	//		log.Println(err)
//	//		return nil, errs.ErrDb
//	//	}
//	//	var exam Exam
//	//	err = json.Unmarshal(key.([]byte), &exam)
//	//	courseDb.Exams = append(courseDb.Exams, exam)
//	//}
//	return &courseDb, nil
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
