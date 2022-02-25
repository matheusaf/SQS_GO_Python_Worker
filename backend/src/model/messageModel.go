package model

// Model representando o objeto a ser recebido na requisicao e o binding a ser realizado para o parse do body da requisicao
type MessageModel struct {
	Channel string `json:"channel" binding:"required"`
	Message string `json:"message" binding:"required"`
}
