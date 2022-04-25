package errs

import "errors"

var (
	ErrRedisNil                = errors.New("redis: nil")
	ErrDb                      = errors.New("unexpected database error")
	ErrUnmarshallingJson       = errors.New("can't unmarshal the JSON to instance")
	ErrDuplicateExam           = errors.New("an exam with this ID exists")
	ErrExamDoesNotExist        = errors.New("this exam doesn't exist")
	ErrQuestionDoesNotExist    = errors.New("this question doesn't exist")
	ErrCourseDoesNotExist      = errors.New("this course doesn't exist")
	ErrExamUpdateId            = errors.New("can't update the exam information because you've changed th exam id")
	ErrRequiredFieldsAreMissed = errors.New("some required fields are missed")
	ErrDuplicateUserExam       = errors.New("this student took this exam before")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
