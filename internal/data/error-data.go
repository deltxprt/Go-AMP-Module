package data

type ErrorMessage struct {
	Title      string `json:"Title"`
	Message    string `json:"Message"`
	StackTrace string `json:"StackTrace"`
}
