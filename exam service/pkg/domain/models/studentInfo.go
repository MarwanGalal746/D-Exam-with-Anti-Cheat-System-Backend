package models

type StudentInfo struct {
	Grade          string `json:"grade" validate:"required,numeric"`
	CheatingStatus string `json:"cheatingStatus" validate:"required,alpha"`
	Report         string `json:"report" validate:"required,alpha"`
}

type StudentGradeRepository interface {
	Add(string, string, string, StudentInfo) error
}
