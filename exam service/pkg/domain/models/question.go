package models

type Question struct {
	Id          string   `json:"id,omitempty"`
	Question    string   `json:"question" validate:"required"`
	RightChoice string   `json:"rightChoice" validate:"required"`
	Choices     []string `json:"choices" validate:"required"`
}
