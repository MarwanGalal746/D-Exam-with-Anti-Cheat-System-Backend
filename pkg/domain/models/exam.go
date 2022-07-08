package models

type ExamInfo struct {
	ExamId          string   `json:"examId" validate:"required"`
	Name            string   `json:"name" validate:"required"`
	CourseId        string   `json:"courseId" validate:"required"`
	TeacherName     string   `json:"teacherName" validate:"required"`
	Duration        int64    `json:"duration" validate:"required"`
	TotalMark       int      `json:"totalMark" validate:"required"`
	Date            int64    `json:"date" validate:"required"`
	QuestionIds     []string `json:"questionIds,omitempty"`
	BlockedStudents []string `json:"blockedStudents"`
}

type Exam struct {
	ExamData  ExamInfo   `json:"examData" validate:"required"`
	Questions []Question `json:"questions,omitempty"`
}

func ResetExamInfo(source ExamInfo) ExamInfo {
	return ExamInfo{ExamId: source.ExamId, Name: source.Name, CourseId: source.CourseId, TeacherName: source.TeacherName, Duration: source.Duration,
		TotalMark: source.TotalMark, Date: source.Date, BlockedStudents: source.BlockedStudents}
}

type ExamRepository interface {
	Create(Exam) error
	GetCourseExams([]string) ([]CourseExams, error)
	GetExam(string, string) (*Exam, error)
	DelExam(string) error
	DelCourseExams(string) error
	UpdateExamInfo(string, ExamInfo) error
}
