package models

type ExamSubmission struct {
	ExamId         string  `json:"examId" validate:"required"`
	CourseId       string  `json:"courseId" validate:"required"`
	UserId         string  `json:"userId" validate:"required"`
	CheatingStatus string  `json:"cheatingStatus" validate:"required"`
	Report         string  `json:"report" validate:"required"`
	Answers        []QsAns `json:"answers" validate:"required"`
}

type QsAns struct {
	QsId   string `json:"questionId" validate:"required"`
	Answer string `json:"answer" validate:"required"`
}

type StudentScore struct {
	Score float64 `json:"score"`
	Total int     `json:"total"`
}

type ExamSubmissionRepository interface {
	SubmitExam(ExamSubmission) (*StudentScore, error)
}
