package exam

type Question struct {
	Id          string   `json:"id,omitempty"`
	Question    string   `json:"question"`
	RightChoice string   `json:"rightChoice"`
	Choices     []string `json:"choices"`
}
