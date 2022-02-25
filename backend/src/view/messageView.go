package view

// View representando o objeto a mensagem a ser enviada caso nao ocorra qualquer tipo de erro
type MessageView struct {
	MessageId         string `json:"messageId"`
	MessageBody       string `json:"messageBody"`
	MessageAttributes string `json:"messageAttributes"`
}
