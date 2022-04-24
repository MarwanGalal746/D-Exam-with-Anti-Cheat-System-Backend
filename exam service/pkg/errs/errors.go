package errs

import "errors"

var (
	ErrRedisNil            = errors.New("redis: nil")
	ErrDb                  = errors.New("unexpected database error")
	ErrUnmarshallingJson   = errors.New("can't unmarshal the JSON to instance")
	ErrMarshallingInstance = errors.New("can't marshal the instance to JSON")
	ErrDuplicateExam       = errors.New("an exam with this ID exists")
	ErrExamDoesNotExist    = errors.New("this exam doesn't exist")
	ErrCourseDoesNotExist  = errors.New("this course doesn't exist")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}
