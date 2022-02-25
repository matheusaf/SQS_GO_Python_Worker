package view

type MessageView struct {
	MessageId         string `json:"messageId"`
	MessageBody       string `json:"messageBody"`
	MessageAttributes string `json:"messageAttributes"`
}
