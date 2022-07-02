package models

type Question struct {
	Id          string   `json:"id,omitempty"`
	Question    string   `json:"question" validate:"required"`
	RightChoice string   `json:"rightChoice" validate:"required"`
	Choices     []string `json:"choices" validate:"required"`
}

func ResetQuestionInfo(source Question) Question {
	return Question{Question: source.Question, RightChoice: source.RightChoice, Choices: source.Choices}
}

type QuestionRepository interface {
	Add(string, Question) (*Question, error)
	Delete(string, string) error
	Update(string, string, Question) error
}
