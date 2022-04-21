package question

type Question struct {
	Id          string   `json:"id"`
	Question    string   `json:"question"`
	RightChoice string   `json:"rightChoice"`
	Choices     []string `json:"choices"`
}
