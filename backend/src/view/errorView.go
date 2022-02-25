package view

// View representando o objeto contendo o erro e sua mensagem ao ser enviada caso ocorra algum erro no envio da mensagem
type ErrorMessage struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
